package permissions

// Type represents a Permission's type.
type Type int

const (
	TypeRole Type = iota + 1
	TypeUser
)
