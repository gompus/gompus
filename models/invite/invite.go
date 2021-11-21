package invite

import (
	"github.com/gompus/gompus/models/application"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/models/guild/event"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/timestamp"
)

// Invite represents an invite to a guild or group DM channel.
// See https://discord.com/developers/docs/resources/invite#invite-resource.
type Invite struct {
	// Code is the invite code.
	Code string `json:"code,omitempty"`

	// Guild is the guild the invite is for.
	Guild guild.Guild `json:"guild,omitempty"`

	// Channel is the channel the invite is for.
	Channel channel.Channel `json:"channel,omitempty"`

	// Inviter is the user who created the invite.
	Inviter user.User `json:"inviter,omitempty"`

	// TargetType is the type of target (for voice channel invites).
	TargetType TargetType `json:"target_type,omitempty"`

	// TargetUser is the user whose stream to
	// display (for voice channel system invites).
	TargetUser user.User `json:"target_user,omitempty"`

	// TargetApplication is the embedded application to
	// open (for voice channel embedded application invites).
	TargetApplication application.Application `json:"target_application,omitempty"`

	// ApproximatePresenceCount is an approximation
	// of the number of online members.
	ApproximatePresenceCount int `json:"approximate_presence_count,omitempty"`

	// ApproximateMemberCount is an approximation
	// of the amount of members in the guild.
	ApproximateMemberCount int `json:"approximate_member_count,omitempty"`

	// ExpiresAt is the invite's expiration date.
	ExpiresAt *timestamp.Timestamp `json:"expires_at,omitempty"`

	// StageInstance contains data about public stage
	// instance (if the invite is for a stage channel).
	StageInstance StageInstance `json:"stage_instance,omitempty"`

	// GuildScheduledEvent contains data about a scheduled event.
	GuildScheduledEvent event.ScheduledEvent `json:"guild_scheduled_event,omitempty"`
}
