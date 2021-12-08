package user

import (
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
)

// GetConnections retrieves a set of connections.
func GetConnections(token auth.Token) (conns []user.Connection, err error) {
	return conns, client.GET(client.Request{
		Path:   "/users/@me/connections",
		Token:  token,
		Result: &conns,
	})
}
