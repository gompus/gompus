package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/invite"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetInvites retrieves a set of invites
// for the channel with the given id.
func GetInvites(token auth.Token, channelID snowflake.Snowflake) (invites []invite.Invite, err error) {
	return invites, client.GET(client.Request{
		Path:   fmt.Sprintf("/channels/%s/invites", channelID),
		Token:  token,
		Result: &invites,
	})
}
