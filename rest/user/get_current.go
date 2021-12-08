package user

import (
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
)

// GetCurrent retrieves the current user.
func GetCurrent(token auth.Token) (user user.User, err error) {
	return user, client.GET(client.Request{
		Path:  "/users/@me",
		Token: token,
	})
}
