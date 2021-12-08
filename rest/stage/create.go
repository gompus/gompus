package stage

import (
	"github.com/gompus/gompus/models/guild/event"
	"github.com/gompus/gompus/models/stage"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type CreateParams struct {
	// ChannelID is the id of the stage channel.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// Topic is the state's topic.
	Topic string `json:"topic,omitempty"`

	// PrivacyLevel is the stage's privacy level (default: GUILD_ONLY).
	PrivacyLevel event.PrivacyLevel `json:"privacy_level,omitempty"`
}

// Create creates a new stage.
func Create(token auth.Token, params CreateParams) (stage stage.Stage, err error) {
	return stage, client.POST(client.Request{
		Path:   "/stage-instances",
		Body:   params,
		Token:  token,
		Result: &stage,
	})
}
