package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// WeChatMessage defines the structure for a WeChat Work message
type WeChatMessage struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

// SendWeChatNotification sends a notification to a WeChat Work webhook
func SendWeComNotification(webhookURL string, content string) error {
	// Create the message payload
	message := WeChatMessage{
		MsgType: "text",
	}
	message.Text.Content = content

	// Convert the message struct to JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Send the POST request
	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the status code indicates success
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send WeChat notification, status code: %d", resp.StatusCode)
	}

	return nil
}
