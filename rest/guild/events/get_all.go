package events

import (
	"fmt"
	"github.com/gompus/gompus/models/guild/event"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type GetAllParams struct {
	// WithUserCount indicates whether to include
	// the number of users subscribed to each event.
	WithUserCount bool `json:"withUserCount,omitempty"`
}

// GetAll retrieves a set of scheduled
// events for the guild with the given id.
func GetAll(token auth.Token, guildID snowflake.Snowflake, params ...GetAllParams) (events []event.ScheduledEvent, err error) {
	return events, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/scheduled-events", guildID),
		Query:  client.GenerateQuery(params),
		Token:  token,
		Result: &events,
	})
}
