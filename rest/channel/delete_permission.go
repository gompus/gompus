package channel

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// DeletePermission deletes the permission overwrite with
// the given id from the channel with the given id.
func DeletePermission(token auth.Token, channelID, overwriteID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/channels/%s/permissions/%s", channelID, overwriteID),
		Token: token,
	})
}
