package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type GetPruneCountParams struct {
	// Days is the number of days to count prune for.
	Days int `json:"days,omitempty"`

	// IncludeRoles contains the roles to include.
	IncludeRoles []snowflake.Snowflake `json:"include_roles,omitempty"`
}

// GetPruneCount retrieves the number of members
// that would be removed in a prune operation
func GetPruneCount(token auth.Token, guildID snowflake.Snowflake, params ...GetPruneCountParams) (int, error) {
	resp := struct {
		Pruned int `json:"pruned,omitempty"`
	}{}
	err := client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/prune", guildID),
		Query:  client.GenerateQuery(params),
		Token:  token,
		Result: &resp,
	})
	return resp.Pruned, err
}
