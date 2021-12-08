package thread

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type GetActiveResponse struct {
	// Threads contains active threads.
	Threads []channel.Channel `json:"threads,omitempty"`

	// Members contains a thread member
	// for each of the returned threads.
	Members []channel.ThreadMember `json:"members,omitempty"`

	// HasMore indicates whether there are potentially
	// additional threads that could be returned with
	// a subsequent request.
	HasMore bool `json:"has_more,omitempty"`
}

// GetActive retrieves a set of active
// threads in the channel with the given id.
func GetActive(token auth.Token, threadID snowflake.Snowflake) (resp GetActiveResponse, err error) {
	return resp, client.GET(client.Request{
		Path:   fmt.Sprintf("/channels/%s/threads/active", threadID),
		Token:  token,
		Result: &resp,
	})
}
