package events

import (
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/snowflake"
)

// TypingStartEvent is the event sent
// when a user starts typing in a channel.
type TypingStartEvent struct {
	// ChannelID is the id of the channel the user started typing in.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// GuildID is the id of the channel's guild.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// UserID is the id of the user who started typing.
	UserID snowflake.Snowflake `json:"user_id,omitempty"`

	// Timestamp is the unix time of when the user started typing.
	Timestamp int `json:"timestamp,omitempty"`

	// Member is the member who started typing
	// if the event occurred in a guild.
	Member guild.Member `json:"member,omitempty"`
}
