package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/channel/message"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// CrosspostMessage crossposts the message with the given id
// in the news channel with the given id to channels following
// the news channel.
func CrosspostMessage(token auth.Token, channelID, messageID snowflake.Snowflake) (msg message.Message, err error) {
	return msg, client.POST(client.Request{
		Path:   fmt.Sprintf("/channels/%s/messages/%s/crosspost", channelID, messageID),
		Token:  token,
		Result: &msg,
	})
}
