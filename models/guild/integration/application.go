package integration

import (
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
)

// Application represents an Integration application.
// See https://discord.com/developers/docs/resources/guild#integration-application-object.
type Application struct {
	// ID is the application's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Name is the application's name.
	Name string `json:"name,omitempty"`

	// Icon is the application's icon hash.
	Icon *string `json:"icon,omitempty"`

	// Description is the application's description.
	Description string `json:"description,omitempty"`

	// Summary is the application's summary.
	Summary string `json:"summary,omitempty"`

	// Bot is the bot associated with the application.
	Bot user.User `json:"bot,omitempty"`
}
