package invite

//go:generate stringer -type=TargetType -trimprefix=TargetType

// TargetType is the target type of an invite.
type TargetType int

const (
	TargetTypeStream TargetType = iota + 1
	TargetTypeEmbeddedApplication
)
