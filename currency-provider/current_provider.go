package currency_provider

import (
	"encoding/json"
)

const (
	endpoint  = "http://www.apilayer.net/api/live?access_key="
	accessKey = "c5ff7182c8795fbb41efb204e4a1a476"
)

type CurrencyProvider interface {
	GetCurrency() (*Currency, error)
}

type networkCurrencyProvider struct {
	requester JSONRequester
}

func (c *networkCurrencyProvider) GetCurrency() (*Currency, error) {
	url := endpoint + accessKey
	body, err := c.requester.Get(url)
	if err != nil {
		return nil, err
	}
	currency := &Currency{}
	err = json.Unmarshal(body, &currency)
	if err != nil {
		return nil, err
	}
	return currency, nil
}

func NewCurrencyProvider() CurrencyProvider {
	return &networkCurrencyProvider{
		requester: DefaultRequester(),
	}
}
