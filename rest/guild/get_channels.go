package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetChannels retrieves a set of channels
// from the guild with the given id.
func GetChannels(token auth.Token, guildID snowflake.Snowflake) (channels []channel.Channel, err error) {
	return channels, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/channels", guildID),
		Token:  token,
		Result: &channels,
	})
}
