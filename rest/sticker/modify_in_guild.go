package sticker

import (
	"fmt"
	"github.com/gompus/gompus/models/sticker"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type ModifyParams struct {
	Name        string  `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Tags        string  `json:"tags,omitempty"`
}

// ModifyInGuild modifies the sticker with the
// given id in the guild with the given id.
func ModifyInGuild(token auth.Token, guildID, stickerID snowflake.Snowflake, params ModifyParams) (sticker sticker.Sticker, err error) {
	return sticker, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/stickers/%s", guildID, stickerID),
		Body:   params,
		Token:  token,
		Result: &sticker,
	})
}
