package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type ModifyPositionsParams struct {
	// ID is the channel's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Position is the channel's position.
	Position *int `json:"position,omitempty"`

	// LockPermissions indicates whether to sync the
	// permission overwrites with the new parent, if
	// moving to a new category.
	LockPermissions *bool `json:"lock_permissions,omitempty"`

	// ParenID is the id of the channel's nwe parent category.
	ParentID *snowflake.Snowflake `json:"parent_id,omitempty"`
}

// ModifyPositions modifies the positions of a set
// of channels in the guild with the given id.
func ModifyPositions(token auth.Token, guildID snowflake.Snowflake, params ModifyPositionsParams) error {
	return client.PATCH(client.Request{
		Path:  fmt.Sprintf("/guilds/%s/channels", guildID),
		Body:  params,
		Token: token,
	})
}
