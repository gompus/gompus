package voice

import (
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

// GuildMember represents the member of a guild.
// See https://discord.com/developers/docs/resources/guild#guild-member-object.
type GuildMember struct {
	// User is the user the member represents.
	User user.User `json:"user,omitempty"`

	// Nick is the user's guild nickname.
	Nick *string `json:"nick,omitempty"`

	// Avatar is the member's guild avatar hash.
	Avatar *string `json:"avatar,omitempty"`

	// Roles contains the ids of the member's roles in the guild.
	Roles []snowflake.Snowflake `json:"roles,omitempty"`

	// JoinedAt is the time when the user joined the guild.
	JoinedAt timestamp.Timestamp `json:"joined_at,omitempty"`

	// PremiumSince is the time when the user started boosting the guild.
	PremiumSince *timestamp.Timestamp `json:"premium_since,omitempty"`

	// Deaf indicates whether the user is deafened in voice channels.
	Deaf bool `json:"deaf,omitempty"`

	// Mute indicates whether the user is muted in voice channels.
	Mute bool `json:"mute,omitempty"`

	// Pending indicates whether the user has no yet passed
	// the guild's membership screening requirements.
	Pending bool `json:"pending,omitempty"`

	// Permissions contains the member's permissions.
	Permissions string `json:"permissions,omitempty"`
}
