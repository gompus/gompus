package channel

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// PinMessage pins the message with the given
// id in the channel with the given id.
func PinMessage(token auth.Token, channelID, messageID snowflake.Snowflake) error {
	return client.PUT(client.Request{
		Path:  fmt.Sprintf("/channels/%s/pin/%s", channelID, messageID),
		Token: token,
	})
}
