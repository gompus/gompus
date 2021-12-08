package user

import (
	"fmt"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Get retrieves the user with the given id.
func Get(token auth.Token, userID snowflake.Snowflake) (user user.User, err error) {
	return user, client.GET(client.Request{
		Path:   fmt.Sprintf("/users/%s", userID),
		Token:  token,
		Result: &user,
	})
}
