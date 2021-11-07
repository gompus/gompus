package team

import "github.com/gompus/snowflake"

// A Team represents a Discord team.
// See https://discord.com/developers/docs/topics/teams#teams.
type Team struct {
	// Icon is the hash of the image of the team's icon.
	Icon *string `json:"icon,omitempty"`

	// ID is the team's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Members is a list of members in the team.
	Members []Member `json:"members,omitempty"`

	// Name is the team's name.
	Name string `json:"name,omitempty"`

	// OwnerUserID is the user id of the team's owner.
	OwnerUserID snowflake.Snowflake `json:"owner_user_id,omitempty"`
}
