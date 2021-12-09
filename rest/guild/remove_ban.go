package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// RemoveBan removes the ban for the user with
// the given id from the guild with the given id.
func RemoveBan(token auth.Token, guildID, userID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/bans/%s", guildID, userID),
		Token: token,
	})
}
