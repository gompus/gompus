package thread

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetMember retrieves the thread member with
// the given id from the thread with the given id.
func GetMember(token auth.Token, threadID, userID snowflake.Snowflake) (member channel.ThreadMember, err error) {
	return member, client.GET(client.Request{
		Path:   fmt.Sprintf("/channels/%s/thread-members/%s", threadID, userID),
		Token:  token,
		Result: &member,
	})
}

// GetAllMembers retrieves a set of thread
// members from the thread with the given id.
func GetAllMembers(token auth.Token, threadID snowflake.Snowflake) (members []channel.ThreadMember, err error) {
	return members, client.GET(client.Request{
		Path:   fmt.Sprintf("/channels/%s/thread-members", threadID),
		Token:  token,
		Result: &members,
	})
}
