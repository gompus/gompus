package thread

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

type GetPrivateArchivedParams struct {
	// Before is the maximum id to include in the results.
	Before timestamp.Timestamp `json:"before,omitempty"`

	// Limit is the maximum number of threads to return.
	Limit int `json:"limit,omitempty"`
}

type GetPrivateArchivedResponse struct {
	// Threads contains the returned threads.
	Threads []channel.Channel `json:"threads,omitempty"`

	// Members contains a thread member
	// for each returned thread.
	Members []channel.ThreadMember `json:"members,omitempty"`

	// HasMore indicates whether there are potentially
	// additional threads that could be returned on a
	// subsequent request.
	HasMore bool `json:"has_more,omitempty"`
}

// GetPrivateArchived retrieves a set of private
// archived threads from the channel with the given id.
func GetPrivateArchived(token auth.Token, channelID snowflake.Snowflake, params ...GetPrivateArchivedParams) (resp GetPrivateArchivedResponse, err error) {
	return resp, client.GET(client.Request{
		Path:  fmt.Sprintf("/channels/%s/threads/archived/private", channelID),
		Query: client.GenerateQuery(params),
		Token: token,
	})
}
