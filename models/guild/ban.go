package guild

import "github.com/gompus/gompus/models/user"

// Ban represents a guild ban.
// See https://discord.com/developers/docs/resources/guild#ban-object.
type Ban struct {
	// Reason is the reason for the ban.
	Reason *string `json:"reason,omitempty"`

	// User is the banned user.
	User user.User `json:"user,omitempty"`
}
