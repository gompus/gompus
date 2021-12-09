package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// AddMemberRole adds the role with the given id
// to the user with the given id in the guild with
// the given id.
func AddMemberRole(token auth.Token, guildID, userID, roleID snowflake.Snowflake) error {
	return client.PUT(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/members/%s/roles/%s", guildID, userID, roleID),
		Token: token,
	})
}
