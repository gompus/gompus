package user

import (
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type CreateGroupDMParams struct {
	// AccessTokens is the set of access tokens of users
	// that have granted the app the gdm.join scope.
	AccessTokens []string `json:"access_tokens,omitempty"`

	// Nicks maps from a user's id to their nickname.
	Nicks map[snowflake.Snowflake]string `json:"nicks,omitempty"`
}

// CreateGroupDM creates a new group
// DM channel with multiple users.
func CreateGroupDM(token auth.Token, params CreateGroupDMParams) (channel channel.Channel, err error) {
	return channel, client.POST(client.Request{
		Path:   "/users/@me/channels",
		Body:   params,
		Token:  token,
		Result: &channel,
	})
}
