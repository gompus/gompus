package team

import (
	"github.com/gompus/snowflake"
)

// A Member represents a user belonging to a Team.
type Member struct {
	// MembershipState is the user's MembershipState.
	MembershipState MembershipState `json:"membership_state,omitempty"`

	// Permissions will always be ["*"] (for whatever reason).
	Permissions []string `json:"permissions,omitempty"`

	// TeamID is the id of the Team the user belongs to.
	TeamID snowflake.Snowflake `json:"team_id,omitempty"`

	// User represents the user.
	User PartialUser `json:"user,omitempty"`
}

type PartialUser struct {
	ID            snowflake.Snowflake `json:"id,omitempty"`
	Username      string              `json:"username,omitempty"`
	Discriminator string              `json:"discriminator,omitempty"`
	Avatar        *string             `json:"avatar,omitempty"`
}
