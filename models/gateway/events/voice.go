package events

import (
	"github.com/gompus/gompus/models/voice"
	"github.com/gompus/snowflake"
)

// VoiceServerUpdateEvent is the event sent when a guild's
// voice server is updated (e.g. when initially connecting
// to voice, and when the current voice instance fails over
// to a new server).
type VoiceServerUpdateEvent struct {
	// Token is the token for the voice connection.
	Token string `json:"token,omitempty"`

	// GuildID is the id of the guild the update is for.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Endpoint represents the voice server host.
	Endpoint *string `json:"endpoint,omitempty"`
}

// VoiceStateUpdateEvent is the event sent when
// someone joins/leaves/moves voice channels.
type VoiceStateUpdateEvent voice.State
