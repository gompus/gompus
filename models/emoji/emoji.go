package emoji

import (
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
)

// An Emoji represents a Discord emoji.
// See https://discord.com/developers/docs/resources/emoji#emoji-object.
type Emoji struct {
	// ID is the emoji's id.
	ID *snowflake.Snowflake `json:"id,omitempty"`

	// Name is the emoji's name.
	Name *string `json:"name,omitempty"`

	// Roles contains the roles allowed to use this emoji.
	Roles []snowflake.Snowflake `json:"roles,omitempty"`

	// User is the user that created this emoji.
	User user.User `json:"user,omitempty"`

	// RequireColons indicates whether this emoji must be wrapped in colons.
	RequireColons bool `json:"require_colons,omitempty"`

	// Managed indicates whether this emoji is managed.
	Managed bool `json:"managed,omitempty"`

	// Animated indicates whether this emoji is animated.
	Animated bool `json:"animated,omitempty"`

	// Available indicates whether this emoji can be used.
	// This may be false due to loss of Server boosts.
	Available bool `json:"available,omitempty"`
}
