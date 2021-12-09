package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type BeginPruneParams struct {
	// Days is the number of days to prune.
	Days int `json:"days,omitempty"`

	// ComputePruneCount indicates whether to
	// include a prune count in the response.
	ComputePruneCount bool `json:"compute_prune_count,omitempty"`

	// IncludeRoles contains the roles
	// to include in the response.
	IncludeRoles []snowflake.Snowflake `json:"include_roles,omitempty"`
}

// BeginPrune begins a prune operation
// in the guild with the given id.
func BeginPrune(token auth.Token, guildID snowflake.Snowflake, params BeginPruneParams) error {
	return client.POST(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/prune", guildID),
		Body:  params,
		Token: token,
	})
}
