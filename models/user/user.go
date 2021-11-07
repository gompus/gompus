package user

import "github.com/gompus/snowflake"

// User represents a Discord user.
// See https://discord.com/developers/docs/resources/user#users-resource.
type User struct {
	// ID is the user's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Username is the user's name.
	Username string `json:"username,omitempty"`

	// Discriminator is the user's 4-digit discord-tag.
	Discriminator string `json:"discriminator,omitempty"`

	// Avatar is the user's avatar hash.
	Avatar *string `json:"avatar,omitempty"`

	// Bot reports whether the user belongs to an OAuth2 application.
	Bot bool `json:"bot,omitempty"`

	// MFAEnabled reports whether the user has two-factor authentication enabled.
	MfaEnabled bool `json:"mfa_enabled,omitempty"`

	// Banner is the user's banner hash.
	Banner *string `json:"banner,omitempty"`

	// AccentColor is the user's banner color
	AccentColor *int `json:"accent_color,omitempty"`

	// Locale is the user's chosen language option.
	Locale string `json:"locale,omitempty"`

	// Verified reports whether the user's email has been verified.
	Verified bool `json:"verified,omitempty"`

	// Email is the user's email.
	Email string `json:"email,omitempty"`

	// Flags is the bitwise or the user's Flags.
	Flags Flags `json:"flags,omitempty"`

	// PremiumType is the type of the Nitro subscription on the user's account.
	PremiumType PremiumType `json:"premium_type,omitempty"`

	// PublicFlags is the bitwise or of public Flags on the user's account.
	PublicFlags Flags `json:"public_flags,omitempty"`
}
