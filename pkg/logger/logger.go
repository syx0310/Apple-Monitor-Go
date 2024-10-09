package logger

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/syx0310/Apple-Monitor-Go/pkg/notify"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func InitLogger() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func ParseJSONResponse(deviceName string, data map[string]interface{}, pushconfig PushConf, storeWhitelist []string) {
	bodyData, ok := data["body"].(map[string]interface{})
	if !ok {
		Error.Printf("%s: Failed to parse 'body' from JSON response\n", deviceName)
		return
	}

	content, ok := bodyData["content"].(map[string]interface{})
	if !ok {
		Error.Printf("%s: Failed to parse 'content' from JSON response\n", deviceName)
		return
	}

	pickupMessage, ok := content["pickupMessage"].(map[string]interface{})
	if !ok {
		Error.Printf("%s: Failed to parse 'pickupMessage' from JSON response\n", deviceName)
		return
	}

	stores, ok := pickupMessage["stores"].([]interface{})
	if !ok {
		Error.Printf("%s: Failed to parse 'stores' from JSON response\n", deviceName)
		return
	}

	availableStores := []string{}
	unavailableStores := []string{}

	for _, store := range stores {
		storeMap, ok := store.(map[string]interface{})
		if !ok {
			continue
		}

		storeName, ok := storeMap["storeName"].(string)
		if !ok {
			continue
		}

		partsAvailability, ok := storeMap["partsAvailability"].(map[string]interface{})
		if !ok {
			continue
		}

		if len(storeWhitelist) > 0 {
			matched := false
			for _, keyword := range storeWhitelist {
				if strings.Contains(storeName, keyword) {
					matched = true
					break
				}
			}
			if !matched {
				// Skip this store as it doesn't match any keyword in the whitelist
				continue
			}
		}

		isAvailable := false
		for _, part := range partsAvailability {
			partMap, ok := part.(map[string]interface{})
			if !ok {
				continue
			}

			pickupDisplay, ok := partMap["pickupDisplay"].(string)
			if !ok {
				continue
			}

			if pickupDisplay == "available" {
				isAvailable = true
				break
			}
		}

		if isAvailable {
			availableStores = append(availableStores, storeName)
		} else {
			unavailableStores = append(unavailableStores, storeName)
		}
	}

	if len(availableStores) > 0 {
		Info.Printf("%s is **available** in the following stores: %v\n", deviceName, availableStores)

		// Send Bark notification
		if pushconfig.BarkKey != "" && pushconfig.BarkAPIURL != "" {
			title := fmt.Sprintf("%s available", deviceName)
			body := fmt.Sprintf("available stores:\n%s", strings.Join(availableStores, "\n"))
			err := notify.SendBarkNotification(pushconfig.BarkKey, pushconfig.BarkAPIURL, title, body)
			if err != nil {
				Error.Printf("%s: Failed to send Bark notification: %v\n", deviceName, err)
			} else {
				Info.Printf("%s: Bark notification sent successfully.\n", deviceName)
			}
		} else if pushconfig.WeComURL != "" {
			// Send WeCom notification
			message := fmt.Sprintf("%s available\navailable stores:\n%s", deviceName, strings.Join(availableStores, "\n"))
			err := notify.SendWeComNotification(pushconfig.WeComURL, message)
			if err != nil {
				Error.Printf("%s: Failed to send WeCom notification: %v\n", deviceName, err)
			} else {
				Info.Printf("%s: WeCom notification sent successfully.\n", deviceName)
			}
		} else {
			Info.Printf("%s: No notification method configured.\n", deviceName)
		}
	} else {
		Info.Printf("%s is **unavailable** in nearby stores (%d stores): %v\n", deviceName, len(unavailableStores), unavailableStores)
	}
}
