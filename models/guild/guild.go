package guild

import (
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/models/emoji"
	"github.com/gompus/gompus/models/gateway"
	"github.com/gompus/gompus/models/gateway/commands"
	"github.com/gompus/gompus/models/guild/welcomescreen"
	"github.com/gompus/gompus/models/permissions"
	"github.com/gompus/gompus/models/stage"
	"github.com/gompus/gompus/models/sticker"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/gompus/models/voice"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

// Guild represents a Discord guild.
// See https://discord.com/developers/docs/resources/guild#guild-resource.
type Guild struct {
	// ID is the guild's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Name is the guild's name.
	Name string `json:"name,omitempty"`

	// Icon is the guild's icon hash.
	Icon *string `json:"icon,omitempty"`

	// IconHash is also the guild's icon hash.
	IconHash *string `json:"icon_hash,omitempty"`

	// Splash is the guild's splash hash.
	Splash *string `json:"splash,omitempty"`

	// SplashHash is also the guild's discovery splash hash.
	DiscoverySplash *string `json:"discovery_splash,omitempty"`

	// Owner indicates whether the current user is the guild's owner.
	Owner bool `json:"owner,omitempty"`

	// OwnerID is the id of the guild's owner.
	OwnerID snowflake.Snowflake `json:"owner_id,omitempty"`

	// Permissions contains the current user's permissions in the guild.
	Permissions string `json:"permissions,omitempty"`

	// AfkChannelID is the id of the guild's afk channel.
	AfkChannelID *snowflake.Snowflake `json:"afk_channel_id,omitempty"`

	// AfkTimeout is the afk timeout, in seconds.
	AfkTimeout int `json:"afk_timeout,omitempty"`

	// WidgetEnabled indicates whether the server widget is enabled.
	WidgetEnabled bool `json:"widget_enabled,omitempty"`

	// WidgetChannelID is the channel id that
	// the widget will generate an invite to.
	WidgetChannelID *snowflake.Snowflake `json:"widget_channel_id,omitempty"`

	// VerificationLevel is the verification level required for the guild.
	VerificationLevel VerificationLevel `json:"verification_level,omitempty"`

	// DefaultMsgNotifications is the default message notification level.
	DefaultMsgNotifications DefaultMsgNotificationLvl `json:"default_msg_notifications,omitempty"`

	// ExplicitContentFilter is the explicit content filter level.
	ExplicitContentFilter ExplicitContentFilterLvl `json:"explicit_content_filter,omitempty"`

	// Roles contains the guild's roles.
	Roles []permissions.Role `json:"roles,omitempty"`

	// Emojis contains the guild's emojis.
	Emojis []emoji.Emoji `json:"emojis,omitempty"`

	// Features contains the guild's enabled features.
	Features []Feature `json:"features,omitempty"`

	// MfaLevel is the guild's required mfa level.
	MfaLevel MfaLevel `json:"mfa_level,omitempty"`

	// ApplicationId is the application id of the guild
	// creator, if the guild was created by a bot.
	ApplicationID *snowflake.Snowflake `json:"application_id,omitempty"`

	// SystemChannelID is the id of the channel where guild
	// notices such as welcome messages and boost events are
	// posted.
	SystemChannelID *snowflake.Snowflake `json:"system_channel_id,omitempty"`

	// SystemChannelFlags contains system channel flags.
	SystemChannelFlags SystemChannelFlags `json:"system_channel_flags,omitempty"`

	// RulesChannelID is the id of the channel where
	// community guilds can display rules and/or guidelines.
	RulesChannelID *snowflake.Snowflake `json:"rules_channel_id,omitempty"`

	// JoinedAt is the time at which the current user joined the guild.
	JoinedAt timestamp.Timestamp `json:"joined_at,omitempty"`

	// Large indicates whether the guild is considered large.
	Large bool `json:"large,omitempty"`

	// Unavailable indicates whether the
	// guild is unavailable due to an outage.
	Unavailable bool `json:"unavailable,omitempty"`

	// MemberCount is the total amount of members in the guild.
	MemberCount int `json:"member_count,omitempty"`

	// VoiceStates contains the states of members
	// currently in voice channels in the guild.
	VoiceStates []voice.State `json:"voice_states,omitempty"`

	// Members contains the guild's members.
	Members []Member `json:"members,omitempty"`

	// Channels contains the guild's channels.
	Channels []channel.Channel `json:"channels,omitempty"`

	// Threads contains active guild threads that
	// the current user has permission to view.
	Threads []channel.Channel `json:"threads,omitempty"`

	// Presences contains presences of the guild's members.
	Presences []PresenceUpdate `json:"presences,omitempty"`

	// MaxPresences is the maximum number of presences for the guild.
	MaxPresences *int `json:"max_presences,omitempty"`

	// MaxMembers is the maximum number of members for the guild.
	MaxMembers int `json:"max_members,omitempty"`

	// VanityUrlCode is the vanity url code for the guild.
	VanityUrlCode *string `json:"vanity_url_code,omitempty"`

	// Description is the description of a community guild.
	Description *string `json:"description,omitempty"`

	// Banner is the guild's banner hash.
	Banner *string `json:"banner,omitempty"`

	// PremiumTier is the guild's boost level.
	PremiumTier PremiumTier `json:"premium_tier,omitempty"`

	// PremiumSubscriptionCount is the number
	// of boosts the guild currently has.
	PremiumSubscriptionCount int `json:"premium_subscription_count,omitempty"`

	// PreferredLocale is the preferred locale of a Community guild.
	PreferredLocale string `json:"preferred_locale,omitempty"`

	// PublicUpdatesChannelID is the id of the channel
	// where admins and moderators of Community guilds
	// receive notices from Discord.
	PublicUpdatesChannelID *snowflake.Snowflake `json:"public_updates_channel_id,omitempty"`

	// MaxVideoChannelUsers is the maximum
	// amount of users in a video channel.
	MaxVideoChannelUsers int `json:"max_video_channel_users,omitempty"`

	// ApproximateMemberCount is an approximation
	// of the amount of members in the guild.
	ApproximateMemberCount int `json:"approximate_member_count,omitempty"`

	// ApproximatePresenceCount is an approximation of
	// the amount of non-offline members in the guild.
	ApproximatePresenceCount int `json:"approximate_presence_count,omitempty"`

	// WelcomeScreen is the welcome screen of a Community guild.
	WelcomeScreen welcomescreen.WelcomeScreen `json:"welcome_screen,omitempty"`

	// NsfwLevel is the guild's nsfw level.
	NsfwLevel NsfwLevel `json:"nsfw_level,omitempty"`

	// StageInstance contains the guild's stage instances.
	StageInstances []stage.Stage `json:"stage_instances,omitempty"`

	// Stickers contains the guild's custom stickers.
	Stickers []sticker.Sticker `json:"stickers,omitempty"`
}

// PresenceUpdate is the event sent when a user's presence
// or info, such as their name or avatar, is updated.
type PresenceUpdate struct {
	// User is the user whose presence is being updated.
	User user.User `json:"user,omitempty"`

	// GuildID is the of the user's guild.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Status is the user's new status.
	Status commands.Status `json:"status,omitempty"`

	// Activites contains the user's current activities.
	Activites []gateway.Activity `json:"activites,omitempty"`

	// ClientStatus is the user's platform-dependent status.
	ClientStatus ClientStatus `json:"client_status,omitempty"`
}

// ClientStatus represents a user's platform-dependent status.
type ClientStatus struct {
	// Desktop is the user's status set for
	// an active desktop application session.
	Desktop string `json:"desktop,omitempty"`

	// Mobile is the user's status set for
	// an active mobile application session.
	Mobile string `json:"mobile,omitempty"`

	// Web is the user's status set for
	// an active web application session.
	Web string `json:"web,omitempty"`
}
