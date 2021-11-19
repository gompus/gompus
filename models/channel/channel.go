package channel

import (
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

// A Channel represents a Discord channel.
// See https://discord.com/developers/docs/resources/channel#channel-object.
type Channel struct {
	// ID is the channel's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Type is the channel's type.
	Type Type `json:"type,omitempty"`

	// GuildID is the id of the guild this channel belongs to.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Position is the position of the channel in the channel list.
	Position int `json:"position,omitempty"`

	// PermissionOverwrites is the set of explicit
	// permission overwrites for members and roles.
	PermissionOverwrites []Overwrite `json:"permission_overwrites,omitempty"`

	// Name is the channel's name.
	Name string `json:"name,omitempty"`

	// Topic is the channel's topic.
	Topic *string `json:"topic,omitempty"`

	// NSFW indicates whether the channel is NSFW.
	NSFW bool `json:"nsfw,omitempty"`

	// LastMessageID is the id of the last message sent in the channel.
	LastMessageID *snowflake.Snowflake `json:"last_message_id,omitempty"`

	// Bitrate is the channel's bitrate (in bits, for voice channels).
	Bitrate int `json:"bitrate,omitempty"`

	// UserLimit is the channel's user limit
	UserLimit int `json:"user_limit,omitempty"`

	// RateLimitPerUser is the amount of seconds a user
	// has to wait before sending another message.
	RateLimitPerUser int `json:"rate_limit_per_user,omitempty"`

	// Recipients is the set of recipients of the DM.
	Recipients []user.User `json:"recipients,omitempty"`

	// Icon is the channel's icon hash.
	Icon *string `json:"icon,omitempty"`

	// OwnerID is the id of the creator of the gruop DM or thread.
	OwnerID snowflake.Snowflake `json:"owner_id,omitempty"`

	// ApplicationID is the application id of the
	// group DM creator (if it is bot-created).
	ApplicationID snowflake.Snowflake `json:"application_id,omitempty"`

	// ParentID is the id of the channel's parent
	// category (only valid for guild channels).
	// For threads, ParentID is the id of the text
	// channel in which the thread was created.
	ParentID *snowflake.Snowflake `json:"parent_id,omitempty"`

	// LastPinTimestamp is the time when
	// the last pinned message was pinned.
	LastPinTimestamp *timestamp.Timestamp `json:"last_pin_timestamp,omitempty"`

	// RTCRegion is the voice region id for voice channels.
	RTCRegion *string `json:"rtc_region,omitempty"`

	// VideoQualityMode is the camera video quality mode for voice channels.
	VideoQualityMode VideoQualityMode `json:"video_quality_mode,omitempty"`

	// MessageCount is an approximate count of messages in a
	// thread. Note that no more than 50 messages are counted.
	MessageCount int `json:"message_count,omitempty"`

	// MemberCount is an approximate count of users in a
	// thread. Note that no more than 50 users are counted.
	MemberCount int `json:"member_count,omitempty"`

	// ThreadMetadata contains thread-specific
	// fields not needed by other channels.
	ThreadMetadata ThreadMetadata `json:"thread_metadata,omitempty"`

	// Member represents the current user, if they have joined the thread.
	Member ThreadMember `json:"member,omitempty"`

	// DefaultAutoArchiveDuration is default duration that the
	// clients will use for newly created threads, in minutes,
	// to automatically archive th thread after recent activity.
	DefaultAutoArchiveduration int `json:"default_auto_archiveduration,omitempty"`

	// Permissions contains computed permissions for the
	// invoking user in the channel, including overwrites.
	Permissions string `json:"permissions,omitempty"`
}
