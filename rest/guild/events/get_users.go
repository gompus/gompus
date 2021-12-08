package events

import (
	"fmt"
	"github.com/gompus/gompus/models/guild/event"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type GetUserParams struct {
	// Limit is the maximum number of users to retrieve.
	Limit int `json:"limit,omitempty"`

	// WithMember indicates whether to include
	// guild member data in the response.
	WithMember bool `json:"with_member,omitempty"`

	// Before is the maximum user id to include in the response.
	Before snowflake.Snowflake `json:"before,omitempty"`

	// After is the minimum user id to include in the response.
	After snowflake.Snowflake `json:"after,omitempty"`
}

// GetUsers retrieves a set of event users subscribed to the
// event with the given id from the guild with the given id.
func GetUsers(token auth.Token, guildID, eventID snowflake.Snowflake, params ...GetUserParams) (users []event.User, err error) {
	return users, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/scheduled-events/%s/users", guildID, eventID),
		Query:  client.GenerateQuery(params),
		Token:  token,
		Result: &users,
	})
}
