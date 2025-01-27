package main

import (
	"fmt"
	"log"

	"github.com/knovellcloud/go-sdk-example/cat-fact-sdk/api"
)

func main() {
	client := api.NewClient()

	// Fetch a random cat fact
	fact, err := client.GetRandomFact()
	if err != nil {
		log.Fatalf("Error fetching cat fact: %v", err)
	}

	fmt.Printf("Random Cat Fact: %s\n", fact.Text)
}
