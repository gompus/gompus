package welcomescreen

import "github.com/gompus/snowflake"

// WelcomeChannel represents a welcome channel in a guild.
// See https://discord.com/developers/docs/resources/guild#welcome-screen-object-welcome-screen-channel-structure.
type WelcomeChannel struct {
	// ID is the channel's id.
	ID snowflake.Snowflake `json:"channel_id,omitempty"`

	// Description is the channel's description.
	Description string `json:"description,omitempty"`

	// EmojiID is the custom emoji id.
	EmojiID *snowflake.Snowflake `json:"emoji_id,omitempty"`

	// EmojiName is...
	//
	//  ... the emoji name (if custom)
	//  ... the unicode character (if standard)
	//  ... nil (if no emoji is set)
	//
	EmojiName *string `json:"emoji_name,omitempty"`
}
