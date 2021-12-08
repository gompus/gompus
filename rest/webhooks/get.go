package webhooks

import (
	"fmt"
	"github.com/gompus/gompus/models/webhook"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Get retrieves the webhook with the given id.
func Get(token auth.Token, hookID snowflake.Snowflake) (hook webhook.Webhook, err error) {
	return hook, client.GET(client.Request{
		Path:   fmt.Sprintf("/webhooks/%s", hookID),
		Token:  token,
		Result: &hook,
	})
}

// GetWithHookToken retrieves the webhook with
// the given id, using the given webhook token.
func GetWithHookToken(token string, hookID snowflake.Snowflake) (hook webhook.Webhook, err error) {
	return hook, client.GET(client.Request{
		Path:   fmt.Sprintf("/webhooks/%s/%s", hookID, token),
		Result: &hook,
	})
}

// GetForChannel retrieves a set of webhooks
// for the chaneel with the given id.
func GetForChannel(token auth.Token, channelID snowflake.Snowflake) (hooks []webhook.Webhook, err error) {
	return hooks, client.GET(client.Request{
		Path:   fmt.Sprintf("/channels/%s/webhooks", channelID),
		Token:  token,
		Result: &hooks,
	})
}

// GetForGuild retrieves a set of webhooks
// for the guild with the given id.
func GetForGuild(token auth.Token, guildID snowflake.Snowflake) (hooks []webhook.Webhook, err error) {
	return hooks, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/webhooks", guildID),
		Token:  token,
		Result: &hooks,
	})
}
