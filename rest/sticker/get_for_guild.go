package sticker

import (
	"fmt"
	"github.com/gompus/gompus/models/sticker"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetAllForGuild retrieves a set of stickers
// for the guild with the given id.
func GetAllForGuild(token auth.Token, guildID snowflake.Snowflake) (stickers []sticker.Sticker, err error) {
	return stickers, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/stickers", guildID),
		Token:  token,
		Result: &stickers,
	})
}

// GetForGuild retrieves the sticker
// with the given id from the given id.
func GetForGuild(token auth.Token, guildID, stickerID snowflake.Snowflake) (sticker sticker.Sticker, err error) {
	return sticker, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/sticker/%s", guildID, stickerID),
		Token:  token,
		Result: &sticker,
	})
}
