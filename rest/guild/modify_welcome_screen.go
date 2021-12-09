package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/guild/welcomescreen"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type ModifyWelcomeScreenParams struct {
	// Description is the server description
	// shown in the welcome screen.
	Description string `json:"description,omitempty"`

	// Enabled indicates whether the
	// welcome screen is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// WelcomeChannels contains channels linked
	// in the welcome screen and their display
	// options.
	WelcomeChannels []welcomescreen.WelcomeChannel `json:"welcome_channels,omitempty"`
}

// ModifyWelcomeScreen modifies the welcome
// screen of the guild with the given id.
func ModifyWelcomeScreen(token auth.Token, guildID snowflake.Snowflake, params ModifyRoleParams) (screen welcomescreen.WelcomeScreen, err error) {
	return screen, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/welcome-screen", guildID),
		Body:   params,
		Token:  token,
		Result: &screen,
	})
}
