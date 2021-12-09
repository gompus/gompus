package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/permissions"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type CreateRoleParams struct {
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
	UniodeEmoji string `json:"uniode_emoji,omitempty"`

	// Mentionable indicates whether
	// the role is mentionable.
	Mentionable bool `json:"mentionable,omitempty"`
}

func CreateRole(token auth.Token, guildID snowflake.Snowflake, params CreateRoleParams) (role permissions.Role, err error) {
	return role, client.POST(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/roles", guildID),
		Token:  token,
		Body:   params,
		Result: &role,
	})
}
