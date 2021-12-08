package templates

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Delete deletes the template with the given
// code from the guild with the given id.
func Delete(token auth.Token, guildID snowflake.Snowflake, code string) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/guild/%s/template/%s", guildID, code),
		Token: token,
	})
}
