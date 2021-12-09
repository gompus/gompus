package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type GetActiveThreadsResponse struct {
	// Threads contains the active threads.
	Threads []channel.Channel `json:"threads,omitempty"`

	// Members contains a thread member for each
	// returned thread the current user has joined.
	Members []channel.ThreadMember `json:"members,omitempty"`
}

// GetActiveThreads retrieves a set of active
// threads in the guild with the given id.
func GetActiveThreads(token auth.Token, guildID snowflake.Snowflake) (resp GetActiveThreadsResponse, err error) {
	return resp, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/threads/active", guildID),
		Token:  token,
		Result: &resp,
	})
}
