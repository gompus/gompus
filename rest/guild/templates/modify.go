package templates

import (
	"fmt"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type ModifyParams struct {
	// Name is the name of the template to be modified.
	Name string `json:"name,omitempty"`

	// Description is the template's description.
	Description string `json:"description,omitempty"`
}

// Modify modifies the metadata of the template with
// the given code in the guild with the given id.
func Modify(token auth.Token, guildID snowflake.Snowflake, code string, params ModifyParams) (template guild.Template, err error) {
	return template, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/guild/%s/template/%s", guildID, code),
		Body:   params,
		Token:  token,
		Result: &template,
	})
}
