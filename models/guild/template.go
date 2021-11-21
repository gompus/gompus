package guild

import (
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

// Template represents a guild template.
// See https://discord.com/developers/docs/resources/guild-template#guild-template-resource.
type Template struct {
	// Code is the template's code.
	Code string `json:"code,omitempty"`

	// Name is the template's name.
	Name string `json:"name,omitempty"`

	// Description is the template's description.
	Description *string `json:"description,omitempty"`

	// UsageCount is the number of times the template has been used.
	UsageCount int `json:"usage_count,omitempty"`

	// CreatorID is the id of the user who created the template.
	CreatorID snowflake.Snowflake `json:"creator_id,omitempty"`

	// Creator is the user who created the template.
	Creator user.User `json:"creator,omitempty"`

	// CreatedAt is the time at which the template was created.
	CreatedAt timestamp.Timestamp `json:"created_at,omitempty"`

	// UpdatedAt is the time at which the template
	// was last synced to the source guild.
	UpdatedAt timestamp.Timestamp `json:"updated_at,omitempty"`

	// SourceGuildID is the id of the guild the template is based on.
	SourceGuildID snowflake.Snowflake `json:"source_guild_id,omitempty"`

	// SerializedSourceGuild represents a snapshot
	// of the guild the template contains.
	SerializedSourceGuild Guild `json:"serialized_source_guild,omitempty"`

	// IsDirty indicates whether the template has unsynced changes.
	IsDirty *bool `json:"is_dirty,omitempty"`
}
