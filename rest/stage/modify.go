package stage

import (
	"fmt"
	"github.com/gompus/gompus/models/stage"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type ModifyParams struct {
	// Topic is the stage's topic.
	Topic string `json:"topic,omitempty"`

	// PrivacyLevel is the stage's privacy level.
	PrivacyLevel stage.PrivacyLevel `json:"privacy_level,omitempty"`
}

// Modify updates attributes of the
// stage in the channel withthe given id.
func Modify(token auth.Token, channelID snowflake.Snowflake, params ModifyParams) (stage stage.Stage, err error) {
	return stage, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/stage-instances/%s", channelID),
		Body:   params,
		Result: &stage,
		Token:  token,
	})
}
