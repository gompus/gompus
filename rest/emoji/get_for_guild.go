package emoji

import (
	"fmt"
	"github.com/gompus/gompus/models/emoji"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetAllForGuild retrieves all emoji for the guild with the given ID.
func GetAllForGuild(token auth.Token, guildID snowflake.Snowflake) (emoji []emoji.Emoji, err error) {
	return emoji, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/emojis", guildID),
		Token:  token,
		Result: &emoji,
	})
}

// GetForGuild retrieves the emoji with the
// given id from the guild with the given id.
func GetForGuild(token auth.Token, guildID, emojiID snowflake.Snowflake) (emoji emoji.Emoji, err error) {
	return emoji, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/emojis/%s", guildID, emojiID),
		Token:  token,
		Result: &emoji,
	})
}
