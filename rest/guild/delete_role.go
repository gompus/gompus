package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// DeleteRole deletes the role with the given
// id from the guild with the given id.
func DeleteRole(token auth.Token, guildID, roleID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/roles/%s", guildID, roleID),
		Token: token,
	})
}
