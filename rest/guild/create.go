package guild

import (
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/models/permissions"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type CreateParams struct {
	// Name is the guild's name.
	Name string `json:"name,omitempty"`

	// Icon is the guild's icon.
	Icon []byte `json:"icon,omitempty"`

	// VerificationLevel is the guild's verification level.
	VerificationLevel guild.VerificationLevel `json:"verification_level,omitempty"`

	// DefaultMessageNotifications is the guild's
	// default message notification level.
	DefaultMsgNotificationLvl guild.DefaultMsgNotificationLvl `json:"default_message_notifications,omitempty"`

	// ExplicitContentFilter is the guild's
	// explicit content filter level.
	ExplicitContentFilter guild.ExplicitContentFilterLvl `json:"explicit_content_filter,omitempty"`

	// Roles contains new guild roles.
	Roles []permissions.Role `json:"roles,omitempty"`

	// Channels contains new guild channels.
	Channels []channel.Channel `json:"channels,omitempty"`

	// AfkChannelID is the afk channel's id.
	AfkChannelID snowflake.Snowflake `json:"afk_channel_id,omitempty"`

	// AfkTimeout is the afk timeout (in seconds).
	AfkTimeout int `json:"afk_timeout,omitempty"`

	// SystemChannelID is the id of the channel where guild
	// notices such as welcome messages and boost events are
	// posted.
	SystemChannelID snowflake.Snowflake `json:"system_channel_id,omitempty"`

	// SystemChannelFlags contains the guild's system channel flags.
	SystemChannelFlags guild.SystemChannelFlags `json:"system_channel_flags,omitempty"`
}

// Create creates a new guild.
func Create(token auth.Token, params CreateParams) (guild guild.Guild, err error) {
	return guild, client.POST(client.Request{
		Path:  "/guilds",
		Body:  params,
		Token: token,
	})
}
