package thread

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Leave removes the current user from
// the thread with the given id.
func Leave(token auth.Token, threadID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/channels/%s/thread-members/@me", threadID),
		Token: token,
	})
}
