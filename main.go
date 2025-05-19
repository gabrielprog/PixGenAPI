package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Environment variables loaded successfully")

	// Example usage of GeneratePix and IsApprovedPayment functions
	// value := 50
	// response := requests.GeneratePix(&value)
	// log.Println("Generated Pix Code:", *response)
	// isApproved := requests.IsApprovedPayment("transaction_id")
	// log.Println("Is payment approved:", isApproved)
	// Uncomment the above lines to test the functions
	// Note: Make sure to replace "transaction_id" with an actual transaction ID for testing
	// Also, ensure that the GeneratePix function is called with a valid value
	// and that the IsApprovedPayment function is called with a valid transaction ID
	// You can also add more test cases or examples as needed
}
