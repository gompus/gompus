package guild

//go:generate stringer -type=DefaultMsgNotificationLvl

// DefaultMsgNotificationLvl represents a guild's default message notification level.
type DefaultMsgNotificationLvl int

const (
	AllMessages  DefaultMsgNotificationLvl = 0
	OnlyMentions                           = 1
)
