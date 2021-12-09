package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/invite"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetVanityUrl retrieves an invite
// for the guild with the given id.
func GetVanityUrl(token auth.Token, guildID snowflake.Snowflake) (invite invite.Invite, err error) {
	return invite, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/vanity-url", guildID),
		Token:  token,
		Result: &invite,
	})
}
