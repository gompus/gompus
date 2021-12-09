package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type AddMemberParams struct {
	// AccessToken is the oauth2 access token granted
	// with the guild.join to the bot's application for
	// the user to add to the guild.
	AccessToken string `json:"access_token,omitempty"`

	// Nick is the nickname to user for the user.
	Nick string `json:"nick,omitempty"`

	// Roles contains ids of roles to assign to the user.
	Roles []snowflake.Snowflake `json:"roles,omitempty"`

	// Mute indicates whether the user
	// is muted in voice channels.
	Mute bool `json:"mute,omitempty"`

	// Deaf indicates whether the user
	// is deafened in voice channels.
	Deaf bool `json:"deaf,omitempty"`
}

// AddMember adds the user with the given
// id to the guild with the given id.
func AddMember(token auth.Token, guildID, userID snowflake.Snowflake, params AddMemberParams) error {
	return client.PUT(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/members/%s", guildID, userID),
		Body:  params,
		Token: token,
	})
}
