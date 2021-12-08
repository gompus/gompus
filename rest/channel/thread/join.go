package thread

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Join adds the current user to
// the thread with the given id.
func Join(token auth.Token, channelID snowflake.Snowflake) error {
	return client.PUT(client.Request{
		Path:  fmt.Sprintf("/channels/%s/thread-members/@me", channelID),
		Token: token,
	})
}
