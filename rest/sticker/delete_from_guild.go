package sticker

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// DeleteFromGuild deletes the sticker with the
// given id from the guild with the given id.
func DeleteFromGuild(token auth.Token, guildID, stickerID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/stickers/%s", guildID, stickerID),
		Token: token,
	})
}
