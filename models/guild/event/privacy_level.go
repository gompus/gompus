package event

//go:generate stringer -type=PrivacyLevel -trimprefix=PrivacyLevel

// PrivacyLevel represents and event's privacy level.
type PrivacyLevel int

// PrivacyLevelGuildOnly is the privacy level
// for events only accessible by guild members.
const PrivacyLevelGuildOnly PrivacyLevel = 2
