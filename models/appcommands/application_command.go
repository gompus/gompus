package appcommands

import "github.com/gompus/snowflake"

// ApplicationCommand is a discord application command.
// See https://discord.com/developers/docs/interactions/application-commands#application-commands.
type ApplicationCommand struct {
	// ID is the command's unique id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Type is the command's type. It defaults to TypeChatInput.
	Type Type `json:"type,omitempty"`

	// ApplicationID is the unique id of the command's parent application.
	ApplicationID snowflake.Snowflake `json:"application_id,omitempty"`

	// GuildID is the id of the guild the command is registered for.
	// If the command is a global command, GuildID will be snowflake.Zero.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Name is the command's name.
	Name string `json:"name,omitempty"`

	// Description is the command's description. It will
	// only be set if the command's Type is TypeChatInput.
	Description string `json:"description,omitempty"`

	// Options contains the command's parameters.
	Options []Option `json:"options,omitempty"`

	// DefaultPermission reports whether the command is enabled
	// by default when the command's app is added to a guild.
	DefaultPermission bool `json:"default_permission"`

	// Version is a version identifier, which is
	// updated during substantial record changes.
	Version snowflake.Snowflake `json:"version,omitempty"`
}
