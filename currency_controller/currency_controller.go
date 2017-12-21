package currency_controller

import (
	. "currency/data_store"
	. "currency/currency-provider"
)

type CurrencyController struct {
	provider  CurrencyProvider
	dataStore DataStore
}

func (c*CurrencyController) GetLastCurrency() {
	currency, err := c.provider.GetCurrency()
	if err != nil {
		return
	}
}

func NewCurrencyController(provider CurrencyProvider, dataStore DataStore) CurrencyController {
	return CurrencyController{
		provider:  provider,
		dataStore: dataStore,
	}
}
