package sticker

import "github.com/gompus/snowflake"

// Item contains the smallest amount of data required to render a sticker.
// See https://discord.com/developers/docs/resources/sticker#sticker-item-object.
type Item struct {
	// ID is the sticker's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Name is the sticker's name.
	Name string `json:"name,omitempty"`

	// FormatType is the type of the sticker's format.
	FormatType FormatType `json:"format_type,omitempty"`
}
