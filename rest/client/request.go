package client

import (
	"bytes"
	"encoding/json"
	"github.com/gompus/gompus/rest/client/auth"
	"io"
	"net/http"
	"net/url"
)

// Request represents an HTTP request.
type Request struct {
	// Method is the request's HTTP method.
	Method Method

	// Path is the relative path the request should be dispatched to.
	Path string

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

	// Token is the auth token that should
	// be used when sending the request.
	Token auth.Token
}

func (r Request) makeHttpRequest() (*http.Request, error) {
	body, err := r.makeHttpRequestBody()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(string(r.Method), resolvePath(r.Path), body)
	if err != nil {
		return nil, err
	}

	if r.Token != "" {
		r.Headers = append(r.Headers, AuthHeader(r.Token))
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
