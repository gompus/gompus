package invites

import (
	"fmt"
	"github.com/gompus/gompus/models/invite"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
)

// Get retrieves an invite for the given code.
func Get(token auth.Token, code string) (invite invite.Invite, err error) {
	return invite, client.GET(client.Request{
		Path:   fmt.Sprintf("/invites/%s", code),
		Token:  token,
		Result: &invite,
	})
}
