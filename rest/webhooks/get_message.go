package webhooks

import (
	"fmt"
	"github.com/gompus/gompus/models/channel/message"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/snowflake"
)

// GetMessage retrieves the message from a previously sent webhook.
func GetMessage(token string, hookID, messageID, threadID snowflake.Snowflake) (msg message.Message, err error) {
	return msg, client.GET(client.Request{
		Path: fmt.Sprintf("/webhooks/%s/%s/message/%s", token, hookID, messageID),
		Query: client.GenerateQuery(struct {
			ThreadID snowflake.Snowflake `json:"thread_id"`
		}{
			ThreadID: threadID,
		}),
	})
}
