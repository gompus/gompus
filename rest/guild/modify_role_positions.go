package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/permissions"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type ModifyRolePositionsParams struct {
	// ID is the id of the role to modify.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Position is the role's sorting position.
	Position *int `json:"position,omitempty"`
}

// ModifyRolePositions modifies the positions
// of roles in the guild with the given id.
func ModifyRolePositions(token auth.Token, guildID snowflake.Snowflake, params ModifyMemberParams) (roles []permissions.Role, err error) {
	return roles, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/roles", guildID),
		Body:   params,
		Token:  token,
		Result: &roles,
	})
}
