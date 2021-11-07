package stage

import "github.com/gompus/snowflake"

// A Stage holds information about a live stage.
// See https://discord.com/developers/docs/resources/stage-instance#stage-instance-resource.
type Stage struct {
	// ID is the stage's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// GuildID is the guild id of the associated stage channel.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// ChannelID is the id of the associated stage channel.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// Topic is the stage's topic.
	Topic string `json:"topic,omitempty"`

	// PrivacyLevel is the stage's PrivacyLevel.
	PrivacyLevel PrivacyLevel `json:"privacy_level,omitempty"`

	// DiscoverableDisabled indicates whether Stage Discovery is disabled.
	DiscoverableDisabled bool `json:"discoverable_disabled,omitempty"`
}
