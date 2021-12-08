package user

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// LeaveGuild leaves the guild with the given id.
func LeaveGuild(token auth.Token, guildID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/users/@me/guilds/%s", guildID),
		Token: token,
	})
}
