package voice

import (
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

// A State represents a user's voice connection status.
// See https://discord.com/developers/docs/resources/voice#voice-resource.
type State struct {
	// GuildID is the id of the guild the state is for.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// ChannelID is the id of the channel the user is connected to.
	ChannelID *snowflake.Snowflake `json:"channel_id,omitempty"`

	// UserId is the id of the user the state is for.
	UserID snowflake.Snowflake `json:"user_id,omitempty"`

	// Member is the guild member the state is for.
	Member GuildMember `json:"member,omitempty"`

	// SessionId is the state's session id.
	SessionID string `json:"session_id,omitempty"`

	// Deaf indicates whether the user is deafened by the server.
	Deaf bool `json:"deaf,omitempty"`

	// Mute indicates whether the user is muted by the server.
	Mute bool `json:"mute,omitempty"`

	// SelfDeaf indicates whether the user is locally deafened.
	SelfDeaf bool `json:"self_deaf,omitempty"`

	// SelfMute indicates whether the user is locally muted.
	SelfMute bool `json:"self_mute,omitempty"`

	// SelfStream indicates whether the user is currently streaming.
	SelfStream bool `json:"self_stream,omitempty"`

	// SelfVideo indicates whether the user's camera is enabled.
	SelfVideo bool `json:"self_video,omitempty"`

	// Suppress indicates whether the user is muted by the current user.
	Suppress bool `json:"suppress,omitempty"`

	// RequestToSpeakTimestamp represents the time when the user requested to speak.
	RequestToSpeaktimestamp *timestamp.Timestamp `json:"request_to_speaktimestamp,omitempty"`
}
