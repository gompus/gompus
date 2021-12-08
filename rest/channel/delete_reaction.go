package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/emoji"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// DeleteReaction deletes the reaction with the given emoji
// from the message with the given id in the channel with
// the given id.
func DeleteReaction(token auth.Token, channelID, messageID snowflake.Snowflake, emoji emoji.Emoji) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/@me", channelID, messageID, emoji),
		Token: token,
	})
}

// DeleteUserReaction deletes another user's reaction with
// the given emoji from the message with the given id in the
// channel with the given id.
func DeleteUserReaction(token auth.Token, channelID, messageID, userID snowflake.Snowflake, emoji emoji.Emoji) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/channels/%s/messages/%s/reactions/%s/%s", channelID, messageID, emoji, userID),
		Token: token,
	})
}
