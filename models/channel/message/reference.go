package message

import "github.com/gompus/snowflake"

// A Reference represents a reference to a Message.
// See https://discord.com/developers/docs/resources/channel#message-reference-object.
type Reference struct {
	// MessageID is the id of the originating message.
	MessageID snowflake.Snowflake `json:"message_id,omitempty"`

	// ChannelID is the id of the originating message's channel.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// GuildID is the id of the originating message's guild.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// FailIfNotExists indicates whether to error
	// if the referenced message does not exist.
	FailIfNotExists bool `json:"fail_if_not_exists,omitempty"`
}
