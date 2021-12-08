package thread

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// AddMember adds the user with the given
// id to the thread with the given id.
func AddMember(token auth.Token, threadID, userID snowflake.Snowflake) error {
	return client.PUT(client.Request{
		Path:  fmt.Sprintf("/channels/%s/thread-members/%s", threadID, userID),
		Token: token,
	})
}
