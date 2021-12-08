package user

import (
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type GetGuildsParams struct {
	// Before is the maximum guild id
	// to include in the response.
	Before snowflake.Snowflake `json:"before,omitempty"`

	// After is the minimum guild id
	// to include in the response.
	After snowflake.Snowflake `json:"after,omitempty"`

	// Limit is the maximum number of
	// guilds to include in the response.
	Limit int `json:"limit,omitempty"`
}

// GetGuilds retrieves a set of guilds the current user is a member of.
func GetGuilds(token auth.Token) (guilds []guild.Guild, err error) {
	return guilds, client.GET(client.Request{
		Path:   "/users/@me/guilds",
		Token:  token,
		Result: guilds,
	})
}
