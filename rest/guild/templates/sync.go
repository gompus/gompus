package templates

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Sync syncs the template with the current
// state of the guild with the given id.
func Sync(token auth.Token, guildID snowflake.Snowflake, code string) error {
	return client.PUT(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/templates/%s", guildID, code),
		Token: token,
	})
}
