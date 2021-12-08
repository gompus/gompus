package emoji

import (
	"fmt"
	"github.com/gompus/gompus/models/emoji"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type CreateParams struct {
	// Name is the name of the emoji.
	Name string `json:"name,omitempty"`

	// Image contains the image data for the emoji.
	Image []byte `json:"image,omitempty"`

	// Roles is the set of roles allowed to use the emoji.
	Roles []snowflake.Snowflake `json:"roles,omitempty"`
}

// CreateForGuild creates a new emoji for the guild with the given id.
func CreateForGuild(token auth.Token, guildID snowflake.Snowflake, params ...CreateParams) (emoji emoji.Emoji, err error) {
	return emoji, client.POST(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/emojis", guildID),
		Body:   client.MergeParams(params),
		Token:  token,
		Result: &emoji,
	})
}
