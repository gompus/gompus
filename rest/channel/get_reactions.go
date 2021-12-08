package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/emoji"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type GetReactionsParams struct {
	// After is the id after which to get users.
	After snowflake.Snowflake `json:"after,omitempty"`

	// Limit is the maximum number of users to retrieve.
	Limit int `json:"limit,omitempty"`
}

// GetReactions retrieves a set of users that reacted with
// the givne emoji to the message with the given id in the
// channel with the given id.
func GetReactions(token auth.Token, channelID, messageID snowflake.Snowflake, emoji emoji.Emoji, params ...GetReactionsParams) (users []user.User, err error) {
	return users, client.GET(client.Request{
		Path:   fmt.Sprintf("/channels/%s/messages/%s/reactions/%s", channelID, messageID, emoji),
		Query:  client.GenerateQuery(params),
		Token:  token,
		Result: &users,
	})
}
