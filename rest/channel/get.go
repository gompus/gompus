package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Get retrieves the channel with the given id.
func Get(token auth.Token, channelID snowflake.Snowflake) (channel channel.Channel, err error) {
	return channel, client.GET(client.Request{
		Path:   fmt.Sprintf("/channel/%s", channelID),
		Token:  token,
		Result: &channel,
	})
}
