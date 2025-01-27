package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRandomFact(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"12345","text":"Cats are great!","source":"user"}`))
	}))
	defer server.Close()

	client := &Client{
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}

	fact, err := client.GetRandomFact()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if fact.Text != "Cats are great!" {
		t.Errorf("Expected 'Cats are great!', got '%s'", fact.Text)
	}
}