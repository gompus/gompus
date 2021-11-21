package event

import (
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

// ScheduledEvent represents a scheduled event in a guild.
// See https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event.
type ScheduledEvent struct {
	// ID is the event's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// GuildID is the id of the guild the event belongs to.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// ChannelID is the id of the channel
	// in which the event will be posted.
	ChannelID *snowflake.Snowflake `json:"channel_id,omitempty"`

	// CreatorID is the id of the user that created the event.
	CreatorID *snowflake.Snowflake `json:"creator_id,omitempty"`

	// Name is the event's name.
	Name string `json:"name,omitempty"`

	// Description is the event's description.
	Description string `json:"description,omitempty"`

	// ScheduledStartTime is the time at which the event will start.
	ScheduledStartTime timestamp.Timestamp `json:"scheduled_start_time,omitempty"`

	// ScheduledEndTime is the time at which the event will end.
	ScheduledEndTime *timestamp.Timestamp `json:"scheduled_end_time,omitempty"`

	// PrivacyLevel is the event's privacy level.
	PrivacyLevel PrivacyLevel `json:"privacy_level,omitempty"`

	// Status is the event's status.
	Status Status `json:"status,omitempty"`

	// EntityType is the event's type.
	EntityType EntityType `json:"entity_type,omitempty"`

	// EntityID is the id of an entity associated with the event.
	EntityID *snowflake.Snowflake `json:"entity_id,omitempty"`

	// EntityMetadata contains additional metadata about the event.
	EntityMetadata *EntityMetadata `json:"entity_metadata,omitempty"`

	// Creator is the user that created the event.
	Creator user.User `json:"creator,omitempty"`

	// UserCount is the number of users subscribed to the event.
	UserCount int `json:"user_count,omitempty"`
}
