package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Request represents an HTTP request.
type Request struct {
	// Method is the request's HTTP method.
	Method Method

	// Url is the url the request should be dispatched to.
	Url string

	// Query contains the request's query parameters.
	Query url.Values

	// Headers contains header funcs
	// providing the request's headers.
	Headers []HeaderFunc

	// Body is the request's body.
	Body interface{}

	// Result contains the value the response
	// body should be decoded into.
	Result interface{}
}

func (r Request) makeHttpRequest() (*http.Request, error) {
	body, err := r.makeHttpRequestBody()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(string(r.Method), r.Url, body)
	if err != nil {
		return nil, err
	}

	for _, header := range r.Headers {
		req.Header.Set(header())
	}

	req.URL.RawQuery = r.Query.Encode()

	return req, nil
}

func (r Request) makeHttpRequestBody() (io.Reader, error) {
	if r.Body == nil {
		return nil, nil
	}
	body, err := json.Marshal(r.Body)
	return bytes.NewReader(body), err
}
