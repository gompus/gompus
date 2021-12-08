package sticker

import (
	"fmt"
	"github.com/gompus/gompus/models/sticker"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type CreateParams struct {
	// Name is the name of the sticker to be created.
	Name string `json:"name,omitempty"`

	// Description is the sticker's decsription.
	Description string `json:"description,omitempty"`

	// Tags contains autocomplete tags for the sticker.
	Tags string `json:"tags,omitempty"`

	// File represents the sticker file.
	File []byte `json:"file,omitempty"`
}

// CreateForGuild creates a new sticker for the guild with the given ID.
func CreateForGuild(token auth.Token, guildID snowflake.Snowflake, params CreateParams) (sticker sticker.Sticker, err error) {
	return sticker, client.POST(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/stickers", guildID),
		Body:   client.GenerateForm(params),
		Token:  token,
		Result: &sticker,
	})
}
