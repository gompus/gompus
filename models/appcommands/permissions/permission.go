package permissions

import "github.com/gompus/snowflake"

// Permission allow the client to enable or disable
// commands for specific users or roles within a guild.
// See https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-application-command-permissions-structure.
type Permission struct {
	// ID is the id of the role or user.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Type is the Permission's type.
	Type Type `json:"type,omitempty"`

	// Allowed reports whether the Permission is granted to the role or user.
	Allowed bool `json:"permission,omitempty"`
}
