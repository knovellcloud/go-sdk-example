package models

// Fact represents a single cat fact
type Fact struct {
	ID     string `json:"_id"`
	Text   string `json:"text"`
	Source string `json:"source"`
}