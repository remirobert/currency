package currency_provider_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "currency/currency-provider"
	"github.com/pkg/errors"
	"encoding/json"
)

type testJSONRequester struct {
	body []byte
	err error
}

func (t *testJSONRequester) Get(url string) (body []byte, err error) {
	return t.body, t.err
}

func TestCurrencyProviderGetCurrencyWithError(t *testing.T) {
	testErr := errors.New("error")
	testRequester := &testJSONRequester{
		body:nil,
		err:testErr,
	}
	currency, err := NewNetworkCurrencyProvider(testRequester).GetCurrency()
	assert.Equal(t, err, testErr)
	assert.Nil(t, currency)
}

func TestCurrencyProviderGetCurrencyWithNoBody(t *testing.T) {
	testRequester := &testJSONRequester{
		body:nil,
		err:nil,
	}
	currency, err := NewNetworkCurrencyProvider(testRequester).GetCurrency()
	assert.NotNil(t, err)
	assert.Nil(t, currency)
}

func TestCurrencyProviderGetCurrencyWithInvalidBody(t *testing.T) {
	invalidBytes := []byte(`invalid json`)

	testRequester := &testJSONRequester{
		body:invalidBytes,
		err:nil,
	}
	currency, err := NewNetworkCurrencyProvider(testRequester).GetCurrency()
	assert.NotNil(t, err)
	assert.Nil(t, currency)
}

func TestCurrencyProviderGetCurrencyWithValidBody(t *testing.T) {
	c := Currency{Source:"source"}
	bytes, _ := json.Marshal(c)

	testRequester := &testJSONRequester{
		body:bytes,
		err:nil,
	}
	currency, err := NewNetworkCurrencyProvider(testRequester).GetCurrency()
	assert.Nil(t, err)
	assert.Equal(t, *currency, c)
}

