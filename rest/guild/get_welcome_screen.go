package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/guild/welcomescreen"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetWelcomeScreen retrieves the welcome
// screen for the guild with the given id.
func GetWelcomeScreen(token auth.Token, guildID snowflake.Snowflake) (screen welcomescreen.WelcomeScreen, err error) {
	return screen, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/welcome_screen", guildID),
		Token:  token,
		Result: &screen,
	})
}
