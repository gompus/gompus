package webhooks

import (
	"fmt"
	"github.com/gompus/gompus/models/webhook"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type ModifyParams struct {
	// Name is the webhook's default name.
	Name string `json:"name,omitempty"`

	// Avatar is the default webhook avatar's image.
	Avatar *[]byte `json:"avatar,omitempty"`

	// ChannelID is the id of the new channel
	// the webhook should be moved to.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`
}

// Modify modifies the webhook with the given id.
func Modify(token auth.Token, hookID snowflake.Snowflake, params ModifyParams) (hook webhook.Webhook, err error) {
	return hook, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/webhooks/%s", hookID),
		Body:   params,
		Token:  token,
		Result: &hook,
	})
}

// ModifyWithToken modifies the webhook with
// the given id, using the given webhook token.
func ModifyWithToken(token string, hookID snowflake.Snowflake, params ModifyParams) (hook webhook.Webhook, err error) {
	return hook, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/webhooks/%s/%s", hookID, token),
		Body:   params,
		Result: &hook,
	})
}
