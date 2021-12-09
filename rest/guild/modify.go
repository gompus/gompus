package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type ModifyParams struct {
	// Name is the guild's name.
	Name string `json:"name,omitempty"`

	// Region is the guild's voice region.
	Region *string `json:"region,omitempty"`

	// VerificationLevel is the guild's verification level.
	VerificationLevel *guild.VerificationLevel `json:"verification_level,omitempty"`

	// DefaultMessageNotifications is the guild's default message notifications.
	DefaultMsgNotifications *guild.DefaultMsgNotificationLvl `json:"default_message_notifications,omitempty"`

	// ExplicitContentFilter is the guild's explicit content filter.
	ExplicitContentFilter *guild.ExplicitContentFilterLvl `json:"explicit_content_filter,omitempty"`

	// AfkChannelID is the id of the guild's afk channel.
	AfkChannelID *snowflake.Snowflake `json:"afk_channel_id,omitempty"`

	// AfkTimeout is the guild's afk timeout in seconds.
	AfkTimeout int `json:"afk_timeout,omitempty"`

	// Icon is the guild's icon.
	Icon *[]byte `json:"icon,omitempty"`

	// OwnerID is the id of the guild's owner.
	OwnerID snowflake.Snowflake `json:"owner_id,omitempty"`

	// Splash is the guild's splash.
	Splash *[]byte `json:"splash,omitempty"`

	// DiscoverySplash is the guild's discovery splash.
	DiscoverySplash *[]byte `json:"discovery_splash,omitempty"`

	// Banner is the guild's banner.
	Banner *[]byte `json:"banner,omitempty"`

	// SystemChannelID is the id of the guild's system channel.
	SystemChannelID snowflake.Snowflake `json:"system_channel_id,omitempty"`

	// SystemChannelFlags are the guild's system channel flags.
	SystemChannelFlags guild.SystemChannelFlags `json:"system_channel_flags,omitempty"`

	// RulesChannelID is the id of the guild's rules channel.
	RulesChannelID *snowflake.Snowflake `json:"rules_channel_id,omitempty"`

	// PublicUpdatesChannelID is the id of the guild's public updates channel.
	PublicUpdatesChannelID *snowflake.Snowflake `json:"public_updates_channel_id,omitempty"`

	// PreferredLocale is the guild's preferred locale.
	PreferredLocale *string `json:"preferred_locale,omitempty"`

	// Features contains the guild's enabled features.
	Features []guild.Feature `json:"features,omitempty"`

	// Description is the guild's description.
	Description *string `json:"description,omitempty"`

	// PremiumProgressBarEnable indicates whether the
	// guild's boost progress bar should be enabled.
	PremiumProgressBarEnable bool `json:"premium_progress_bar_enable,omitempty"`
}

// Modify modifies the settings of the guild with the given id.
func Modify(token auth.Token, guildID snowflake.Snowflake, params ModifyParams) (guild guild.Guild, err error) {
	return guild, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/guilds/%s", guildID),
		Body:   params,
		Token:  token,
		Result: &guild,
	})
}
