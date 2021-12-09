package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type GetParams struct {
	// WithCounts indicates whether to include approximate
	// member and presence counts in the result.
	WithCounts bool `json:"with_counts,omitempty"`
}

// Get retrieves the guild with the given id.
func Get(token auth.Token, guildID snowflake.Snowflake, params ...GetParams) (guild guild.Guild, err error) {
	return guild, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s", guildID),
		Query:  client.GenerateQuery(params),
		Token:  token,
		Result: &guild,
	})
}
