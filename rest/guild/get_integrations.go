package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/guild/integration"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetIntegrations retrieves a set of integrations
// from the guild with the given id.
func GetIntegrations(token auth.Token, guildID snowflake.Snowflake) (integrations []integration.Integration, err error) {
	return integrations, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/integrations", guildID),
		Token:  token,
		Result: &integrations,
	})
}
