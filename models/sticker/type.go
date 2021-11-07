package sticker

//go:generate stringer -type=Type

// Type is a Sticker's type.
type Type int

const (
	// Standard is the type for official stickers in a pack
	// , part of Nitro or in a removed purchasable pack.
	Standard Type = 1

	// Guild is the type for stickers uploaded to
	// a boosted guild for the guild's members.
	Guild
)
