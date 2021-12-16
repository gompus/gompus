package events

import (
	"github.com/gompus/gompus/models/application"
	"github.com/gompus/gompus/models/invite"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

// InviteCreateEvent is the event sent when a
// new invite to a guild channel is created.
type InviteCreateEvent struct {
	// ChannelID is the id of the channel the invite is for.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// Code is the invite's code.
	Code string `json:"code,omitempty"`

	// CreatedAt is the time at which the invite was created.
	CreatedAt timestamp.Timestamp `json:"created_at,omitempty"`

	// GuildID is the id of the guild the invite is for.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Inviter is the user that created the invite.
	Inviter user.User `json:"inviter,omitempty"`

	// MaxAge is the duration (in s) after which the invite expires.
	MaxAge int `json:"max_age,omitempty"`

	// MaxUses is the maximum number of times the invite can be used.
	MaxUses int `json:"max_uses,omitempty"`

	// TargetType is the target type for voice channel invites.
	TargetType invite.TargetType `json:"target_type,omitempty"`

	// TargetUser is the user whose stream to
	// display (for voice channel stream invites).
	TargetUser user.User `json:"target_user,omitempty"`

	// TargetApplication is the embedded application
	// to open (for voice channel stream invites).
	TargetApplication application.Application `json:"target_application,omitempty"`

	// Temporary indicates whether the invite is temporary.
	Temporary bool `json:"temporary,omitempty"`

	// Uses is the amount of times the invite has already been used.
	Uses int `json:"uses,omitempty"`
}

// InviteDeleteEvent is the event sent when an invite is deleted.
type InviteDeleteEvent struct {
	// ChannelID is the id of the channel the invite was for.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// GuildID is the id of the guild the invite was for.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Code is the invite's code.
	Code string `json:"code,omitempty"`
}
