package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/channel/message"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetPinnedMessages retrieves a set of pinned
// messages from the channel with the given id.
func GetPinnedMessages(token auth.Token, channelID snowflake.Snowflake) (messages []message.Message, err error) {
	return messages, client.GET(client.Request{
		Path:   fmt.Sprintf("/channels/%s/pinned", channelID),
		Token:  token,
		Result: &messages,
	})
}
