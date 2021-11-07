package permissions

//go:generate stringer -type=Type -trimprefix=Type

// Type represents a Permission's type.
type Type int

const (
	TypeRole Type = iota + 1
	TypeUser
)
