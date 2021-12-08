package channel

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// TriggerTypingIndicator posts a typing
// indicator for channel with the given id.
func TriggerTypingIndicator(token auth.Token, channelID snowflake.Snowflake) error {
	return client.POST(client.Request{
		Path:  fmt.Sprintf("/channels/%s/typing", channelID),
		Token: token,
	})
}
