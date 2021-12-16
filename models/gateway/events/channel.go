package events

import (
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

// ChannelCreateEvent is the event sent
// when a new guild channel is created.
type ChannelCreateEvent channel.Channel

// ChannelDeleteEvent is the event sent when a
// channel relevant to the current user is deleted.
type ChannelDeleteEvent channel.Channel

// ChannelPinsUpdateEvent is the event sent when a
// message is pinned or unpinned in a text channel.
type ChannelPinsUpdateEvent struct {
	// GuildID is the id of the channel's guild.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// ChannelID is the channel's id.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// LastPinTimestamp is the time at which the
	// most recent pinned message was pinned.
	LastPinTimestamp *timestamp.Timestamp `json:"last_pin_timestamp,omitempty"`
}

// ChannelUpdateEvent is the event that is
// sent when a guild channel is updated.
type ChannelUpdateEvent channel.Channel
