package api

import (
	"encoding/json"
	"github.com/knovellcloud/go-sdk-example/cat-fact-sdk/models"
)

// GetRandomFact fetches a random cat fact
func (c *Client) GetRandomFact() (*models.Fact, error) {
	resp, err := c.SendRequest("/facts/random")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var fact models.Fact
	if err := json.NewDecoder(resp.Body).Decode(&fact); err != nil {
		return nil, err
	}

	return &fact, nil
}
