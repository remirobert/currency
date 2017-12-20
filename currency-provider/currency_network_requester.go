package currency_provider

import (
	"net/http"
	"io/ioutil"
)

type JSONRequester interface {
	Get(url string) (body []byte, err error)
}

type NetworkRequester struct {
}

func (n *NetworkRequester) Get(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return
}

func DefaultRequester() JSONRequester {
	return &NetworkRequester{}
}
