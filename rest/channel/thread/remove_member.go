package thread

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// RemoveMember removes the user with the
// given id from the thread with the given id.
func RemoveMember(token auth.Token, threadID, userID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/channels/%s/thread-members/%s", threadID, userID),
		Token: token,
	})
}
