package permissions

import "github.com/gompus/snowflake"

// RoleTags represent a Role's tags.
// See https://discord.com/developers/docs/topics/permissions#role-object-role-tags-structure.
type RoleTags struct {
	// BotID is the id of the role's bot.
	BotID snowflake.Snowflake `json:"bot_id,omitempty"`

	// IntegrationID is the id of this role's integration.
	IntegrationID snowflake.Snowflake `json:"integration_id,omitempty"`

	// PremiumSubscriber indicates whether this
	// is the guild's premium subscriber role.
	PremiumSubscriber bool `json:"premium_subscriber,omitempty"`
}
