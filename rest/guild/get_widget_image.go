package guild

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type WidgetStyle string

const (
	WidgetStyleShield  WidgetStyle = "shield"
	WidgetStyleBanner1 WidgetStyle = "banner1"
	WidgetStyleBanner2 WidgetStyle = "banner2"
	WidgetStyleBanner3 WidgetStyle = "banner3"
	WidgetStyleBanner4 WidgetStyle = "banner4"
)

// GetWidgetImage retrieves a png image
// widget for the guild with the given id.
func GetWidgetImage(token auth.Token, guildID snowflake.Snowflake, style WidgetStyle) (image []byte, err error) {
	return image, client.GET(client.Request{
		Path: fmt.Sprintf("/guilds/%s/widget.png?style=%s", guildID, style),
		Body: struct {
			WidgetStyle WidgetStyle `json:"widget_style,omitempty"`
		}{
			WidgetStyle: style,
		},
		Token:  token,
		Result: &image,
	})
}
