package webhooks

import (
	"fmt"
	"github.com/gompus/gompus/models/webhook"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type CreateParams struct {
	// Name is the name of the webhook to be created.
	Name string `json:"name,omitempty"`

	// Avatar is the image for the default webhook avatar.
	Avatar *[]byte `json:"avatar,omitempty"`
}

// Create creates a new webhook for the channel with the given id.
func Create(token auth.Token, channelID snowflake.Snowflake) (hook webhook.Webhook, err error) {
	return hook, client.POST(client.Request{
		Path:   fmt.Sprintf("/channels/%s/webhooks", channelID),
		Token:  token,
		Result: &hook,
	})
}
