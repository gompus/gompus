package user

import (
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
)

type ModifyParams struct {
	// Username is the user's name.
	Username string `json:"username,omitempty"`

	// Avatar represents the user's avatar.
	Avatar *[]byte `json:"avatar,omitempty"`
}

// ModifyCurrent modifies the current user.
func ModifyCurrent(token auth.Token, params ModifyParams) (user user.User, err error) {
	return user, client.PATCH(client.Request{
		Path:   "/users/@me",
		Body:   params,
		Token:  token,
		Result: &user,
	})
}
