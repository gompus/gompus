package channel

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GroupDMRemoveRecipient removes the user with the
// given id from the group dm channel with the given id.
func GroupDMRemoveRecipient(token auth.Token, channelID, userID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/channels/%s/recipients/%s", channelID, userID),
		Token: token,
	})
}
