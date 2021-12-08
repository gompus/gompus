package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/channel/message"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type GetMessagesParams struct {
	// Around can be used to retrieve
	// messages around some message id.
	Around snowflake.Snowflake `json:"around,omitempty"`

	// Around can be used to retrieve
	// messages before some message id.
	Before snowflake.Snowflake `json:"before,omitempty"`

	// Around can be used to retrieve
	// messages after some message id.
	After snowflake.Snowflake `json:"after,omitempty"`

	// Limit is the maximum number of messages to retrieve.
	Limit int `json:"limit,omitempty"`
}

// GetMessages retrieves a set of messages
// from the channel with the given id.
func GetMessages(token auth.Token, channelID snowflake.Snowflake, params ...GetMessagesParams) (messages []message.Message, err error) {
	return messages, client.GET(client.Request{
		Path:   fmt.Sprintf("/channels/%s/messages", channelID),
		Query:  client.GenerateQuery(params),
		Token:  token,
		Result: &messages,
	})
}
