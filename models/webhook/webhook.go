package webhook

import (
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
)

// A Webhook represents a low-effot way to post messages to Discord channels.
// See https://discord.com/developers/docs/resources/webhook#webhook-resource.
type Webhook struct {
	// Id is the webhook's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Type is the webhook's type.
	Type Type `json:"type,omitempty"`

	// GuildID is the id of the guild the webhook is for.
	GuildID *snowflake.Snowflake `json:"guild_id,omitempty"`

	// ChannelID is the id of the channel the webhook is for.
	ChannelID *snowflake.Snowflake `json:"channel_id,omitempty"`

	// User represents the user who created the webhook.
	User user.User `json:"user,omitempty"`

	// Name is the webhooks' default name.
	Name *string `json:"name,omitempty"`

	// Avatar is the default webhook's user avatar hash.
	Avatar *string `json:"avatar,omitempty"`

	// Token is the webhook's secure token.
	Token string `json:"token,omitempty"`

	// ApplicationID is the id of the bot/OAuth2 application that created the webhook.
	ApplicationID *snowflake.Snowflake `json:"application_id,omitempty"`

	// SourceGuild is the guild of the channel the webhook is following.
	SourceGuild guild.Guild `json:"source_guild,omitempty"`

	// SourceChannel is the channel the webhook is following.
	SourceChannel channel.Channel `json:"source_channel,omitempty"`

	// URL is the url used for executing the webhook.
	URL string `json:"url,omitempty"`
}
