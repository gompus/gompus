package events

import "github.com/gompus/snowflake"

// WebhookUpdateEvent is the event sent when a guild
// channel's webhook is created, updated or deleted.
type WebhookUpdateEvent struct {
	// ChannelID is the id of the channel whose webhook was updated.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// GuildID is the id of the channel's guild.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`
}
