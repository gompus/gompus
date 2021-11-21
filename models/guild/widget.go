package guild

import "github.com/gompus/snowflake"

// Widget represents a guild widget.
// See https://discord.com/developers/docs/resources/guild#guild-widget-object.
type Widget struct {
	// Enabled indicates whether the widget is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// ChannelID is the id of the widget's channel.
	ChannelID *snowflake.Snowflake `json:"channel_id,omitempty"`
}
