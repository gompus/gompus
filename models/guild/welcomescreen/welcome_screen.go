package welcomescreen

// WelcomeScreen represents a welcome screen for a guild.
// See https://discord.com/developers/docs/resources/guild#welcome-screen-object.
type WelcomeScreen struct {
	// Description is the description of
	// the server the welcome screen is for.
	Description *string `json:"description,omitempty"`

	// WelcomeChannels contains the channels shown
	// in the welcome screen (no more than 5).
	WelcomeChannels []WelcomeChannel `json:"welcome_channels,omitempty"`
}
