package user

//go:generate stringer -type=Visibility -trimprefix=Visibility

// Visibility represents a User's visibility.
type Visibility int

const (
	// VisibilityNone indicates the user is invisible to everyone except themselves.
	VisibilityNone Visibility = 0

	// VisibilityEveryone indicates that the user is visible to everyone.
	VisibilityEveryone = 1
)
