package invites

import (
	"fmt"
	"github.com/gompus/gompus/models/invite"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
)

// Delete deletes the invite with the given code.
func Delete(token auth.Token, code string) (invite invite.Invite, err error) {
	return invite, client.DELETE(client.Request{
		Path:   fmt.Sprintf("/invites/%s", code),
		Token:  token,
		Result: &invite,
	})
}
