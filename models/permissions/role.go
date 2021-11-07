package permissions

import "github.com/gompus/snowflake"

// A Role represents a set of permission attached to a group of users.
// See https://discord.com/developers/docs/topics/permissions#role-object.
type Role struct {
	// ID is the role's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Name is the role's name.
	Name string `json:"name,omitempty"`

	// Color is the integer representation of the role's hexadecimal color code.
	Color int `json:"color,omitempty"`

	// Hoist indicates whether the role is pinned in the user listing.
	Hoist bool `json:"hoist,omitempty"`

	// Icon is the role's icon hash.
	Icon *string `json:"icon,omitempty"`

	// UnicodeEmoji is the role's unicode emoji.
	UnicodeEmoji *string `json:"unicode_emoji,omitempty"`

	// Position is the role's position.
	Position int `json:"position,omitempty"`

	// Permissions is the role's permission bit set.
	Permissions string `json:"permissions,omitempty"`

	// Managed indicates whether the roles is managed by an integration.
	Managed bool `json:"managed,omitempty"`

	// Mentionable indicates whether the role is mentionable.
	Mentionable bool `json:"mentionable,omitempty"`

	// Tags contains the role's tags.
	Tags RoleTags `json:"tags,omitempty"`
}
