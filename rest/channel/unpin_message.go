package channel

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// UnpinMessage unpins the message with the
// given id in the channel with the given id.
func UnpinMessage(token auth.Token, channelID, messageID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/channels/%s/pins/%s", channelID, messageID),
		Token: token,
	})
}
