package webhooks

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Delete deletes the webhook with the given id.
func Delete(token auth.Token, hookID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/webhooks/%s", hookID),
		Token: token,
	})
}

// DeleteWithToken deletes the webhook with
// the given id, using the given webhook token.
func DeleteWithToken(token string, hookID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path: fmt.Sprintf("/webhooks/%s/%s", hookID, token),
	})
}
