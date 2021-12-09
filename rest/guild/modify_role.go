package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/permissions"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type ModifyRoleParams struct {
	// Name is the role's name.
	Name string `json:"name,omitempty"`

	// Permissions is the bitwise value
	// of the role's permissions.
	Permissions string `json:"permissions,omitempty"`

	// Color is the role's rgb color value.
	Color int `json:"color,omitempty"`

	// Hoist indicates whether the role should
	// be displayed separately in the sidebar.
	Hoist bool `json:"hoist,omitempty"`

	// Icon is the role's icon.
	Icon []byte `json:"icon,omitempty"`

	// UnicodeEmoji is the role's unicode emoji.
	UnicodeEmoji string `json:"unicode_emoji,omitempty"`

	// Mentionable indicates whether
	// the role is be mentionable.
	Mentionable bool `json:"mentionable,omitempty"`
}

// ModifyRole modifies the role with the
// given id in the guild with the given id.
func ModifyRole(token auth.Token, guildID, roleID snowflake.Snowflake, params ModifyRoleParams) (role permissions.Role, err error) {
	return role, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/roles/%s", guildID, roleID),
		Token:  token,
		Body:   params,
		Result: &role,
	})
}
