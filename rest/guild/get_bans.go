package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetBans retrieves a set of bans for the users
// banned from the guild with the given id.
func GetBans(token auth.Token, guildID snowflake.Snowflake) (bans []guild.Ban, err error) {
	return bans, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/bans", guildID),
		Token:  token,
		Result: &bans,
	})
}

// GetBan retrieves the ban for the user with the
// given id from the guild with the given id.
func GetBan(token auth.Token, guildID, userID snowflake.Snowflake) (ban guild.Ban, err error) {
	return ban, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/bans/%s", guildID, userID),
		Token:  token,
		Result: &ban,
	})
}
