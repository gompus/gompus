package channel

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// DeleteMessage deletes the message with the
// given id from the channel with the given id.
func DeleteMessage(token auth.Token, channelID, messageID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/channels/%s/messages/%s", channelID, messageID),
		Token: token,
	})
}

// BulkDeleteMessages deletes the messages with the
// given ids from the channel with the given id.
func BulkDeleteMessages(token auth.Token, channelID snowflake.Snowflake, messageIDs ...snowflake.Snowflake) error {
	return client.POST(client.Request{
		Path: fmt.Sprintf("/channels/%s/messages/bulk_delete", channelID),
		Body: struct {
			Messages []snowflake.Snowflake `json:"messages"`
		}{
			Messages: messageIDs,
		},
		Token: token,
	})
}
