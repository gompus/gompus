package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/voice"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetVoiceRegions retrieves a set of voice
// regions from the guild with the given id.
func GetVoiceRegions(token auth.Token, guildID snowflake.Snowflake) (regions []voice.Region, err error) {
	return regions, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/regions", guildID),
		Token:  token,
		Result: &regions,
	})
}
