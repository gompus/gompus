package templates

import (
	"fmt"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetAll retrieves a set of guild templates
// for the guild with the given id.
func GetAll(token auth.Token, guildID snowflake.Snowflake) (templates []guild.Template, err error) {
	return templates, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/templates", guildID),
		Token:  token,
		Result: &templates,
	})
}

// Get retrieves the template with the given code.
func Get(token auth.Token, code string) (template guild.Template, err error) {
	return template, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/templates/%s", code),
		Token:  token,
		Result: &template,
	})
}
