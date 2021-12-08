package channel

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type EditPermissionsParams struct {
	// Allow is the bitwise value of all allowed permissions.
	Allow string `json:"allow,omitempty"`

	// Deny is the bitwise value of all denied permissions.
	Deny string `json:"deny,omitempty"`

	// Type indicates the permission's type.
	Type int `json:"type,omitempty"`
}

// EditPermissions edits the permission overwrites
// with the given id for the channel with the given id.
func EditPermissions(token auth.Token, channelID, overwriteID snowflake.Snowflake) error {
	return client.PUT(client.Request{
		Path:  fmt.Sprintf("/channels/%s/permissions/%s", channelID, overwriteID),
		Token: token,
	})
}
