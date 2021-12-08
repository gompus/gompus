package emoji

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gompus/gompus/models/emoji"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type ModifyParams struct {
	// Name is the name of the emoji to modify.
	Name string `json:"name,omitempty"`

	// Roles is the set of roles allowed to use the emoji.
	Roles *[]snowflake.Snowflake `json:"roles,omitempty"`
}

// ModifyInGuild modifies the emoji with the given id in the
// guild with the given id. It returns the updated emoji.
func ModifyInGuild(token auth.Token, guildID snowflake.Snowflake, emojiID snowflake.Snowflake, params ...ModifyParams) (emoji emoji.Emoji, err error) {
	return emoji, client.PATCH(client.Request{
		Path: spew.Sprintf("/guilds/%s/emojis/%s", guildID, emojiID),
		Headers: []client.HeaderFunc{
			client.AuthHeader(token),
		},
		Body:   client.MergeParams(params),
		Result: &emoji,
	})
}
