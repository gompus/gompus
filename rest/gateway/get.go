package gateway

import (
	"github.com/gompus/gompus/rest/client"
)

type GetParams struct {
	// Version is the gateway version to use.
	Version int `json:"v,omitempty"`

	// Encoding is the preferred encoding
	// of date sent over the gateway.
	Encoding Encoding `json:"encoding,omitempty"`

	// Compress can be set to "zlib-stream" to enable
	// zlib compression of data sent over the gateway.
	Compress string `json:"compress,omitempty"`
}

// GetWebsocket retrieves a WWS URL that clients
// can use for connecting to a websocket.
func GetWebsocket(params GetParams) (string, error) {
	var resp struct {
		URL string `json:"url"`
	}
	err := client.GET(client.Request{
		Path:   "/gateway",
		Query:  client.GenerateQuery(params),
		Result: &resp,
	})
	return resp.URL, err
}
