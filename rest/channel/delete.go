package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Delete deletes the channel with the given id. Equivalently,
// Delete closes the private message with the given id.
func Delete(token auth.Token, channelID snowflake.Snowflake) (channel channel.Channel, err error) {
	return channel, client.DELETE(client.Request{
		Path:   fmt.Sprintf("/channel/%s", channelID),
		Result: &channel,
		Token:  token,
	})
}
