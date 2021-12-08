package events

import (
	"fmt"
	"github.com/gompus/gompus/models/guild/event"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type GetParams struct {
	// WithUserCount indicates whether to include
	// the number of users subscribed to the event.
	WithUserCount bool `json:"with_user_count,omitempty"`
}

// Get retrieves the event with the given
// id from the guild with the given id.
func Get(token auth.Token, guildID, eventID snowflake.Snowflake, params ...GetParams) (event event.ScheduledEvent, err error) {
	return event, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/scheduled-events/%s", guildID, eventID),
		Query:  client.GenerateQuery(params),
		Token:  token,
		Result: &event,
	})
}
