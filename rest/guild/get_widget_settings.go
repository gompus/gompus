package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetWidgetSettings retrieves the widget
// for the guild with the given id.
func GetWidgetSettings(token auth.Token, guildID snowflake.Snowflake) (widget guild.Widget, err error) {
	return widget, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/widget", guildID),
		Token:  token,
		Result: &widget,
	})
}
