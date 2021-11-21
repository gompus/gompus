package guild

//go:generate stringer -type=ExplicitContentFilterLevel -timprefix=ContentFilter

// ExplicitContentFilterLvl represents a guild's explicit content filter level.
type ExplicitContentFilterLvl int

const (
	ContentFilterDisabled ExplicitContentFilterLvl = iota
	ContentFilterMembersWithoutRoles
	ContentFilterAllMembers
)
