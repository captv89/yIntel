package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/captv89/yIntel/models"
)

// Get Currency Exchange
func GetExchangeRate(fromCurrency string) models.Currency {
	url := fmt.Sprintf("https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/%s/inr.json", fromCurrency)
	resp, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// Print response body
	var currency models.Currency
	err = json.NewDecoder(resp.Body).Decode(&currency)
	if err != nil {
		log.Fatal(err)
	}
	return currency
}
