package apple

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	cron "github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"github.com/syx0310/Apple-Monitor-Go/pkg/logger"
)

func StartMonitoring() error {
	c := cron.New(cron.WithSeconds())

	for _, device := range AppConfig.Devices {
		device := device

		_, err := c.AddFunc(device.Crontab, func() {
			err := monitorDevice(device)
			if err != nil {
				logger.Error.Printf("Error monitoring device %s: %v\n", device.Name, err)
			}
		})

		if err != nil {
			logger.Error.Printf("Error adding cron job for device %s: %v\n", device.Name, err)
			continue
		}
		logger.Info.Printf("Added cron job for device %s with schedule %s\n", device.Name, device.Crontab)
	}

	c.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	logger.Info.Println("Received interrupt signal, shutting down...")
	c.Stop()
	return nil
}

func monitorDevice(device Device) error {
	regionConfig, err := GetRegionConfig(device.Region)
	if err != nil {
		return fmt.Errorf("failed to get region config for device %s: %v", device.Name, err)
	}

	prefix := regionConfig.Prefix

	queryParams := url.Values{}

	for key, value := range regionConfig.DefaultParams {
		queryParams.Set(key, value)
	}

	for key, value := range device.QueryParams {
		queryParams.Set(key, value)
	}

	queryParams.Set("parts.0", device.ProductID)

	queryParams.Set("location", device.Location)

	fullURL := fmt.Sprintf("%s/shop/fulfillment-messages?%s", prefix, queryParams.Encode())

	if viper.GetBool("verbose") {
		fmt.Printf("Fetching data for device %s from URL: %s\n", device.Name, fullURL)
	}

	resp, err := http.Get(fullURL)
	if err != nil {
		return fmt.Errorf("HTTP request failed for device %s: %v", device.Name, err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body for device %s: %v", device.Name, err)
	}

	if viper.GetBool("verbose") {
		fmt.Printf("Raw JSON response for device %s:\n%s\n", device.Name, string(bodyBytes))
	}

	var data map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return fmt.Errorf("JSON unmarshal failed for device %s: %v", device.Name, err)
	}

	pushConfig := logger.PushConf{
		BarkKey:    device.BarkKey,
		BarkAPIURL: device.BarkAPIURL,
		WeComURL:   device.WeComURL,
	}

	logger.ParseJSONResponse(device.Name, data, pushConfig, device.StoreWhitelistKeyword)

	return nil
}
