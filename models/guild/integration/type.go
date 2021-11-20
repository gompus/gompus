package integration

// Type represents the type of an Integration.
type Type string

const (
	TypeTwitch  Type = "twitch"
	TypeYoutube Type = "youtube"
	TypeDiscord Type = "discord"
)
