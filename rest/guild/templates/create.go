package templates

import (
	"fmt"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type CreateParams struct {
	// Name is the name of the template to be created.
	Name string `json:"name,omitempty"`

	// Description is the template's description.
	Description *string `json:"description,omitempty"`
}

// Create creates a template for the guild with the given id.
func Create(token auth.Token, guildID snowflake.Snowflake, params CreateParams) (template guild.Template, err error) {
	return template, client.POST(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/templates", guildID),
		Body:   params,
		Token:  token,
		Result: &template,
	})
}
