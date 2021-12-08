package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type ModifyGroupDMParams struct {
	// Name is the channel's name.
	Name string `json:"name,omitempty"`

	// Icon is the DM's icon.
	Icon []byte `json:"icon,omitempty"`
}

type ModifyGuildChannelParams struct {
	// Name is the channel's name.
	Name string `json:"name,omitempty"`

	// Type is the channel's type.
	Type channel.Type `json:"type,omitempty"`

	// Position is the channel's position in the left-hand listing.
	Position *int `json:"position,omitempty"`

	// Topic is the channel's topic.
	Topic *string `json:"topic,omitempty"`

	// NSFW indicates whether the channel is NSFW.
	NSFW *bool `json:"nsfw,omitempty"`

	// RateLimit is the amount of seconds a user
	// has to wait before sending another message.
	RateLimit *int `json:"rate_limit_per_user,omitempty"`

	// Bitrate is the channel's bitrate (in bits).
	Bitrate *int `json:"bitrate,omitempty"`

	// UserLimit is the channel's user limit.
	UserLimit *int `json:"user_limit,omitempty"`

	// PermissionOverwrites contains channel
	// or category-specific permissions.
	PermissionOverwrites *[]channel.Overwrite `json:"permission_overwrites,omitempty"`

	// ParentID is the id of the
	// channel's new parent category.
	ParentID *snowflake.Snowflake `json:"parent_id,omitempty"`

	// RTCRegion is the channel's voice region id.
	RTCRegion *string `json:"rtc_region,omitempty"`

	// VideoQualityMode is the channel's
	// camera video quality mode.
	VideoQualityMode *channel.VideoQualityMode `json:"video_quality_mode,omitempty"`

	// DefaultAutoArchiveDuration is the default duration
	// the clients use for newly created threads in the
	// channel, in minutes, to automatically archive the
	// thread after recent activity.
	DefaultAutoArchiveDuration *int `json:"default_auto_archive_duration,omitempty"`
}

type ModifyThreadParams struct {
	// Name is the channel's name.
	Name string `json:"name,omitempty"`

	// Archived indicates whether the thread is archived.
	Archived bool `json:"archived,omitempty"`

	// AutoArchiveDuration is the duration in minutes
	// to automatically archive the thread after recent
	// activity. Possible values are 60, 1440, 4320, 10080.
	AutoArchiveDuration int `json:"auto_archive_duration,omitempty"`

	// Locked indicates whether the thread is locked.
	Locked bool `json:"locked,omitempty"`

	// Invitable indicates whether non-moderators
	// can add other non-moderators to the thread.
	Invitable bool `json:"invitable,omitempty"`

	// RateLimitPerUser is the amount of seconds a
	// user has to wait before sending another message.
	RateLimitPerUser *int `json:"rate_limit_per_user,omitempty"`
}

// Modify modifies the settings of the channel with the
// given id. Params should be of type ModifyGroupDMParams,
// ModifyGuildChannelParams and ModifyThreadParams
func Modify(token auth.Token, channelID snowflake.Snowflake, params ...interface{}) (channel channel.Channel, err error) {
	return channel, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/channels/%s", channelID),
		Body:   client.MergeParams(params),
		Token:  token,
		Result: &channel,
	})
}
