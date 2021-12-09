package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/permissions"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetRoles retrieves a set of roles
// from the guild with the given id.
func GetRoles(token auth.Token, guildID snowflake.Snowflake) (roles []permissions.Role, err error) {
	return roles, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/roles", guildID),
		Token:  token,
		Result: &roles,
	})
}
