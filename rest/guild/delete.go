package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Delete deletes the guild with the given id.
func Delete(token auth.Token, guildID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/guilds/%s", guildID),
		Token: token,
	})
}
