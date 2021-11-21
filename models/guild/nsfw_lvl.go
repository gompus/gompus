package guild

//go:generate stringer type=NsfwLevel

// NsfwLevel represents a guild's nsfw level.
type NsfwLevel int

const (
	DefaultNswfLevel NsfwLevel = iota
	ExplicitNswfLevel
	SafeNsfwLevel
	AgeRestrictedNsfwLevel
)
