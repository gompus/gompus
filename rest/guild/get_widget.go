package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetWidget retrieves the widget
// for the guild with the given id.
func GetWidget(token auth.Token, guildID snowflake.Snowflake) (widget guild.Widget, err error) {
	return widget, client.GET(client.Request{
		Path:   fmt.Sprintf("/guild/%s/widget.json", guildID),
		Token:  token,
		Result: &widget,
	})
}
