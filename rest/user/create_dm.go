package user

import (
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// CreateDM creates a new DM channel
// with the user with the given id.
func CreateDM(token auth.Token, recipientID snowflake.Snowflake) (channel channel.Channel, err error) {
	return channel, client.POST(client.Request{
		Path: "/users/@me/channels",
		Body: struct {
			RecipientID snowflake.Snowflake `json:"recipient_id"`
		}{
			RecipientID: recipientID,
		},
		Token:  token,
		Result: &channel,
	})
}
