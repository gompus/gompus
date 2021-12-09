package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// RemoveMemberRole removes the role with the given
// id from the member with the given id in the guild
// with the given id.
func RemoveMemberRole(token auth.Token, guildID, userID, roleID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/members/%s/roles/%s", guildID, userID, roleID),
		Token: token,
	})
}
