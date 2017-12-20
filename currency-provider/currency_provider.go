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

type NetworkCurrencyProvider struct {
	requester JSONRequester
}

func (c *NetworkCurrencyProvider ) GetCurrency() (*Currency, error) {
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

func NewNetworkCurrencyProvider(requester JSONRequester) *NetworkCurrencyProvider {
	return &NetworkCurrencyProvider {
		requester: requester,
	}
}

func NewCurrencyProvider() CurrencyProvider {
	return &NetworkCurrencyProvider {
		requester: DefaultRequester(),
	}
}
