package guild

//go:generate stringer type=VerificationLevel -trimprefix=VerificationLevel

// VerificationLevel represents a guild's verification level.
type VerificationLevel int

const (
	VerificationLevelNone VerificationLevel = iota
	VerificationLevelLow
	VerificationLevelMedium
	VerificationLevelHigh
	VerificationLevelVeryHigh
)
