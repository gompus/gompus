package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

type ModifyCurrentVoiceStateParams struct {
	// ChannelID is the id of the channel to move to.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// Suppress indicates whether to
	// toggle the user's suppress state.
	Suppress bool `json:"suppress,omitempty"`

	// RequestToSpeakTimestamp sets
	// the user's request to speak.
	RequestToSpeakTimestamp *timestamp.Timestamp `json:"request_to_speak_timestamp,omitempty"`
}

// ModifyCurrentVoiceState modifies the current
// user's voice state in the guild with the given id.
func ModifyCurrentVoiceState(token auth.Token, guildID snowflake.Snowflake, params ModifyCurrentVoiceStateParams) error {
	return client.PATCH(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/voice-states/@me", guildID),
		Body:  params,
		Token: token,
	})
}

type ModifyUserVoiceStateParams struct {
	// ChannelID is the id of the channel to move the user to.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// Suppress indicates whether to
	// toggle the user's suppress state.
	Suppress bool `json:"suppress,omitempty"`
}

// ModifyUserVoiceState modifies the voice state of the
// user with the given id in the guild with the given id.
func ModifyUserVoiceState(token auth.Token, guildID, userID snowflake.Snowflake, params ModifyUserVoiceStateParams) error {
	return client.PATCH(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/voice-states/%s", guildID, userID),
		Body:  params,
		Token: token,
	})
}
