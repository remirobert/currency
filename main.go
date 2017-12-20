package main

import (
	"currency/currency-provider"
	"log"
)

func main() {
	provider := currency_provider.NewCurrencyProvider()
	currency, err := provider.GetCurrency()
	if err != nil {
		log.Print("get error : ", err)
		return
	}
	log.Print(currency)
}
