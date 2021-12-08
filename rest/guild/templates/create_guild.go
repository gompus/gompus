package templates

import (
	"fmt"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
)

type CreateGuildParams struct {
	// Name is the name of the guild to be created.
	Name string `json:"name,omitempty"`

	// Icon is the guild's icon.
	Icon []byte `json:"icon,omitempty"`
}

// CreateGuild creates a new guild based
// on the template with the given code.
func CreateGuild(token auth.Token, code string, params CreateGuildParams) (guild guild.Guild, err error) {
	return guild, client.POST(client.Request{
		Path:   fmt.Sprintf("/guilds/templates/%s", code),
		Body:   params,
		Token:  token,
		Result: &guild,
	})
}
