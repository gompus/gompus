package client

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
)

// GET dispatches the given request as a get request.
func GET(req Request) error {
	req.Method = Get
	return do(req)
}

// POST dispatches the given request as a post request.
func POST(req Request) error {
	req.Method = Post
	return do(req)
}

// PUT dispatches the given request as a put request.
func PUT(req Request) error {
	req.Method = Put
	return do(req)
}

// PATCH dispatches the given request as a patch request.
func PATCH(req Request) error {
	req.Method = Patch
	return do(req)
}

// DELETE dispatches the given request as a delete request.
func DELETE(req Request) error {
	req.Method = Delete
	return do(req)
}

var dialer = &net.Dialer{
	Timeout: 5 * time.Second,
}

var client = &http.Client{
	Transport: &http.Transport{
		DialContext: dialer.DialContext,
	},
	Timeout: 5 * time.Second,
}

func do(req Request) error {
	request, err := req.makeHttpRequest()
	if err != nil {
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(req.Body)
}
