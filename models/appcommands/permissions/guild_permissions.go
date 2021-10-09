package permissions

import "github.com/gompus/snowflake"

// GuildPermissions represents the permissions for a command in a guild.
// See https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-guild-application-command-permissions-structure.
type GuildPermissions struct {
	// ID is the command's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// ApplicationID is the id of the application the command belongs to.
	ApplicationID snowflake.Snowflake `json:"application_id,omitempty"`

	// GuildID is the id of the guild the command belongs to.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Permissions is the set of permissions for the command in the guild.
	Permissions []Permission `json:"permissions,omitempty"`
}
