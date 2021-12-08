package events

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Delete deletes the event with the given
// id from the guild with the given id.
func Delete(token auth.Token, guildID, eventID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/scheduled-events/%s", guildID, eventID),
		Token: token,
	})
}
