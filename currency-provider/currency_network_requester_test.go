package currency_provider_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "currency/currency-provider"
	"github.com/pkg/errors"
	"net/http"
	"bytes"
	"io"
	"encoding/json"
)

type testReaderCloser struct {
	io.Reader
}

func (t *testReaderCloser) Close() error {
	return nil
}

type testNetworkClient struct {
	bytes []byte
	err   error
}

func (t *testNetworkClient) Get(url string) (resp *http.Response, err error) {
	if t.bytes != nil {
		r := bytes.NewReader(t.bytes)
		testReader := &testReaderCloser{r}
		return &http.Response{Body: testReader}, t.err
	}
	return &http.Response{Body: nil}, t.err
}

func TestNetworkRequesterWithError(t *testing.T) {
	testErr := errors.New("error")
	requester := NetworkRequester{Client: &testNetworkClient{err: testErr}}
	bytes, err := requester.Get("")
	assert.NotNil(t, err)
	assert.Equal(t, err, testErr)
	assert.Nil(t, bytes)
}

func TestNetworkRequesterWithNoBody(t *testing.T) {
	testNetClient := &testNetworkClient{bytes: nil}
	requester := NetworkRequester{testNetClient}
	bytes, err := requester.Get("")
	assert.NotNil(t, err)
	assert.Nil(t, bytes)
}

func TestNetworkRequesterWithValidResponse(t *testing.T) {
	c := Currency{Source: "source"}
	bytes, _ := json.Marshal(c)

	testNetClient := &testNetworkClient{bytes: bytes}
	requester := NetworkRequester{Client: testNetClient}
	respBytes, err := requester.Get("")

	assert.Nil(t, err)
	assert.NotNil(t, respBytes)
	assert.Equal(t, respBytes, bytes)
}
