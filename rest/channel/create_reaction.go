package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/emoji"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// CreateReaction creates a reaction for the message with
// the given id in the channel with the given id.
func CreateReaction(token auth.Token, channelID, messageID snowflake.Snowflake, emoji emoji.Emoji) error {
	return client.PUT(client.Request{
		Path:  fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/@me", channelID, messageID, emoji),
		Token: token,
	})
}
