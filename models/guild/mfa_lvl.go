package guild

//go:generate stringer -type=MfaLevel -trimprefix=MfaLevel

// MfaLevel represents a guild's multi-factor-authentication level.
type MfaLevel int

const (
	MfaLevelNone     MfaLevel = 0
	MfaLevelElevated          = 1
)
