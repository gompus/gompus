package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type CreateBanParams struct {
	// DeleteMessageDays is the number
	// of days to delete messages for.
	DeleteMessageDays int `json:"delete_message_days,omitempty"`

	// Reason is the reason for the ban.
	Reason string `json:"reason,omitempty"`
}

// CreateBan creates a ban for the user with
// the given id in the guild with the given id.
func CreateBan(token auth.Token, guildID, userID snowflake.Snowflake, params CreateBanParams) error {
	return client.PUT(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/bans/%s", guildID, userID),
		Body:  params,
		Token: token,
	})
}
