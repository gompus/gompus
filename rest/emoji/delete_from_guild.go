package emoji

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// DeleteFromGuild deletes the emoji with the
// given id from the guild with the given id.
func DeleteFromGuild(token auth.Token, guildID, emojiID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  spew.Sprintf("/guilds/%s/emojis/%s", guildID, emojiID),
		Token: token,
	})
}
