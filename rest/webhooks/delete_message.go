package webhooks

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/snowflake"
)

// DeleteMessage deletes the message with the given id
// that was created by the webhook with the given id.
func DeleteMessage(token string, hookID, messageID, threadID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path: fmt.Sprintf("/webhooks/%s/%s/messages/%s", hookID, token, messageID),
		Query: client.GenerateQuery(struct {
			ThreadID snowflake.Snowflake `json:"thread_id"`
		}{
			ThreadID: threadID,
		}),
	})
}
