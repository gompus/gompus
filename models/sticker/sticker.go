package sticker

import (
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
)

// Sticker represents a Discord sticker that can be sent in messages.
// See https://discord.com/developers/docs/resources/sticker#sticker-resource.
type Sticker struct {
	// ID is the sticker's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// PackID is the pack the sticker is from.
	PackID snowflake.Snowflake `json:"pack_id,omitempty"`

	// Name is the sticker's name.
	Name string `json:"name,omitempty"`

	// Description is the sticker's description.
	Description *string `json:"description,omitempty"`

	// Tags contains autocomplete tags for the sticker. By
	// convention, Tags is a comma-separated list of keywords.
	Tags string `json:"tags,omitempty"`

	// Type is the sticker's type.
	Type Type `json:"type,omitempty"`

	// FormatType is the type of the sticker's format.
	FormatType FormatType `json:"format_type,omitempty"`

	// Available indicates whether the sticker can be
	// used. May be false due to loss of Server boosts.
	Available bool `json:"available,omitempty"`

	// GuildID is the id of the guild that owns the sticker
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// User is the user that uploaded the sticker.
	User user.User `json:"user,omitempty"`

	// SortValue represents the standard sticker's sort order within its pack.
	SortValue int `json:"sort_value,omitempty"`
}
