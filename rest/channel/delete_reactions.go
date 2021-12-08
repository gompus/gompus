package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/emoji"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// DeleteReactions deletes all reactions on the message
// with the given id in the channel with the given id.
func DeleteReactions(token auth.Token, channelID, messageID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/channels/%s/messages/%s/reactions", channelID, messageID),
		Token: token,
	})
}

// DeleteAllReactions deletes all reactions for the given
// emoji from the message with the given id in the channel
// with the given id.
func DeleteAllReactions(token auth.Token, channelID, messageID snowflake.Snowflake, emoji emoji.Emoji) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/channels/%s/messages/%s/reactions/%s", channelID, messageID, emoji),
		Token: token,
	})
}
