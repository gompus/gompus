package invite

import "github.com/gompus/timestamp"

// Metadata contains additional information about an invite.
type Metadata struct {
	// Uses is the number of times the invite has been used already.
	Uses int `json:"uses,omitempty"`

	// MaxUses is the maximum number of times the invite can be used.
	MaxUses int `json:"max_uses,omitempty"`

	// MaxAge is the duration after which the invite expires.
	MaxAge int `json:"max_age,omitempty"`

	// Temporary indicates whether the invite
	// only grants temporary membership.
	Temporary bool `json:"temporary,omitempty"`

	// CreatedAt is the time at which the invite was created.
	CreatedAt timestamp.Timestamp `json:"created_at,omitempty"`
}
