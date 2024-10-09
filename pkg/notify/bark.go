package notify

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SendBarkNotification sends a notification using the Bark app
func SendBarkNotification(barkKey, barkAPIURL, title, body string) error {
	// Construct the URL
	// URL format: {barkAPIURL}/{barkKey}/{title}/{body}
	// Encode the title and body to ensure special characters are handled
	encodedTitle := url.QueryEscape(title)
	encodedBody := url.QueryEscape(body)

	// Remove any trailing slashes from the barkAPIURL
	barkAPIURL = strings.TrimRight(barkAPIURL, "/")

	// Construct the final URL
	notificationURL := fmt.Sprintf("%s/%s/%s/%s", barkAPIURL, barkKey, encodedTitle, encodedBody)

	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a new GET request
	req, err := http.NewRequest("GET", notificationURL, nil)
	if err != nil {
		return err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the status code indicates success
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send Bark notification, status code: %d", resp.StatusCode)
	}

	return nil
}
