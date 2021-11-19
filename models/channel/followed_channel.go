package channel

import "github.com/gompus/snowflake"

// A FollowedChannel represents a channel that a user follows.
// See https://discord.com/developers/docs/resources/channel#followed-channel-object.
type FollowedChannel struct {
	// ChannelID is the id of the source channel.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// WebhookID is the id of the created target webhook.
	WebhookID snowflake.Snowflake `json:"webhook_id,omitempty"`
}
