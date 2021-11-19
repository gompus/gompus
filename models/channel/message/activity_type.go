package message

//go:generate stringer -type=ActivityType -trimprefix=ActivityType

// ActivityType is the type of an Activity.
type ActivityType int

const (
	ActivityTypeJoin        ActivityType = 1
	ActivityTypeSpectate                 = 2
	ActivityTypeListen                   = 3
	ActivityTypeJoinRequest              = 5
)
