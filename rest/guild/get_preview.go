package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetPreview retrieves a guild preview
// for the guild with the given id.
func GetPreview(token auth.Token, guildID snowflake.Snowflake) (preview guild.Preview, err error) {
	return preview, client.GET(client.Request{
		Path:   fmt.Sprintf("/guild/%s/preview", guildID),
		Token:  token,
		Result: &preview,
	})
}
