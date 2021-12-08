package stage

import (
	"fmt"
	"github.com/gompus/gompus/models/stage"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Get retrieves the stage associated
// with the channel with the given id.
func Get(token auth.Token, channelID snowflake.Snowflake) (stage stage.Stage, err error) {
	return stage, client.GET(client.Request{
		Path:   fmt.Sprintf("/stage-instances/%s", channelID),
		Result: &stage,
		Token:  token,
	})
}
