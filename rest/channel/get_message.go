package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/channel/message"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetMessage retrieves the message with the
// given id from the channel with the given id.
func GetMessage(token auth.Token, channelID, messageID snowflake.Snowflake) (msg message.Message, err error) {
	return msg, client.GET(client.Request{
		Path:   fmt.Sprintf("/channel/%s/message/%s", channelID, messageID),
		Token:  token,
		Result: &msg,
	})
}
