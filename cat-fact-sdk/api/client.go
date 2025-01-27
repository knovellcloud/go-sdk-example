package api

import (
	"errors"
	"net/http"
	"time"
)

// Client struct for interacting with the Cat Fact API
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient initializes the SDK client
func NewClient() *Client {
	return &Client{
		BaseURL: "https://cat-fact.herokuapp.com",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// SendRequest handles HTTP requests
func (c *Client) SendRequest(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("non-200 response: " + resp.Status)
	}

	return resp, nil
}
