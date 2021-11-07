package stage

//go:generate stringer -type=PrivacyLvl

// PrivacyLevel determines a Stage's visibility.
type PrivacyLevel int

const (
	// Public is the PrivacyLevel for publicly visible stages.
	Public PrivacyLevel = 1

	// GuildOnly is the PrivacyLevel for stages
	// that are only visible to guild members.
	GuildOnly
)
