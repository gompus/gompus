package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// RemoveMember removes the member with the
// given id from the guild with the given id.
func RemoveMember(token auth.Token, guildID, userID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/members/%s", guildID, userID),
		Token: token,
	})
}
