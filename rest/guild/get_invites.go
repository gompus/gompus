package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/invite"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetInvites retrieves a set of invites
// for the guild with the given id.
func GetInvites(token auth.Token, guildID snowflake.Snowflake) (invites []invite.Invite, err error) {
	return invites, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/invites", guildID),
		Token:  token,
		Result: &invites,
	})
}
