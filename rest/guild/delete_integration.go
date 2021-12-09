package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// DeleteIntegration deletes the integration with
// the given id from the guild with the given id.
func DeleteIntegration(token auth.Token, guildID, integrationID snowflake.Snowflake) error {
	return client.DELETE(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/integrations/%s", guildID, integrationID),
		Token: token,
	})
}
