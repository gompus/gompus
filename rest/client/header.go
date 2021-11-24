package client

import (
	"github.com/gompus/gompus/rest/client/auth"
)

// A HeaderFunc provides a request header.
type HeaderFunc func() (key string, value string)

// AuthHeader returns a HeaderFunc providing an
// authorization header using the given token.
func AuthHeader(token auth.Token) HeaderFunc {
	return func() (key string, value string) {
		return "Authorization", string(token)
	}
}
