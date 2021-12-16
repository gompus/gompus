package commands

import "github.com/gompus/snowflake"

// UpdateVoiceState is the command for clients seeking
// to join, move or disconnect from a voice channel.
// See https://discord.com/developers/docs/topics/gateway#update-voice-state.
type UpdateVoiceState struct {
	// GuildID is the id of the guild.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// ChannelID is the id of the voice channel the client
	// seeks to join (should be nil for disconnecting from
	// a voice channel).
	ChannelID *snowflake.Snowflake `json:"channel_id,omitempty"`

	// SelfMute indicates whether the client
	// should be muted upon joining.
	SelfMute bool `json:"self_mute,omitempty"`

	// SelfDeaf indicates whether the client
	// should be deafened upon joining.
	SelfDeaf bool `json:"self_deaf,omitempty"`
}
