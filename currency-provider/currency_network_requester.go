package currency_provider

import (
	"net/http"
	"io/ioutil"
	"errors"
)

const (
	errorInvalidBody     = "invalid response body"
	errorInvalidResponse = "invalid response"
)

type JSONRequester interface {
	Get(url string) (body []byte, err error)
}

type NetworkClient interface {
	Get(url string) (resp *http.Response, err error)
}

type NetworkRequester struct {
	Client NetworkClient
}

func (n *NetworkRequester) handleResponse(resp *http.Response) (body []byte, err error) {
	if resp.Body == nil {
		return nil, errors.New(errorInvalidBody)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return
}

func (n *NetworkRequester) Get(url string) (body []byte, err error) {
	resp, err := n.Client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New(errorInvalidResponse)
	}
	return n.handleResponse(resp)
}

func DefaultRequester() JSONRequester {
	return &NetworkRequester{
		Client: http.DefaultClient,
	}
}
