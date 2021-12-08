package stage

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Delete deltes the stage from the channel with the given id.
func Delete(token auth.Token, channelID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/stage-instances/%s", channelID),
		Token: token,
	})
}
