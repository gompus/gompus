package guild

import (
	"github.com/gompus/gompus/models/emoji"
	"github.com/gompus/snowflake"
)

// Preview represents a guild preview.
// See https://discord.com/developers/docs/resources/guild#guild-preview-object.
type Preview struct {
	// ID is the guild's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Name is the guild's name.
	Name string `json:"name,omitempty"`

	// Icon is the guild's icon hash.
	Icon *string `json:"icon,omitempty"`

	// Splash is the guild's splash hash.
	Splash *string `json:"splash,omitempty"`

	// DiscoverySplash is the guild's discovery splash hash.
	DiscoverySplash *string `json:"discovery_splash,omitempty"`

	// Emojis contains the guild's emojis.
	Emojis []emoji.Emoji `json:"emojis,omitempty"`

	// Features contains the guild's features.
	Features []Feature `json:"features,omitempty"`

	// ApproximateMemberCount is an approximation
	// of the number of members in the guild.
	ApproximateMemberCount int `json:"approximate_member_count,omitempty"`

	// ApproximatePresenceCount is an approximation
	// of the number of online members in the guild.
	ApproximatePresenceCount int `json:"approximate_presence_count,omitempty"`

	// Description is the guild's description
	// (present if the guild is discoverable).
	Description *string `json:"description,omitempty"`
}
