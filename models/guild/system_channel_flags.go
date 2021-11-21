package guild

import "strings"

// SystemChannelFlags contains guild system channel flags.
type SystemChannelFlags int

const (
	SuppressJoinNotifications SystemChannelFlags = 1 << iota
	SuppressPremiumSubscriptions
	SuppressGuildReminderNotifications
	SuppressJoinNotificationReplies
)

var sysChanFlagsToString = map[SystemChannelFlags]string{
	SuppressJoinNotifications:          "SuppressJoinNotifications",
	SuppressPremiumSubscriptions:       "SuppressPremiumSubscriptions",
	SuppressGuildReminderNotifications: "SuppressGuildReminderNotifications",
	SuppressJoinNotificationReplies:    "SuppressJoinNotificationReplies",
}

// String returns the string representation of flags.
func (flags SystemChannelFlags) String() string {
	var out []string
	for flag, s := range sysChanFlagsToString {
		if flags&flag != 0 {
			out = append(out, s)
		}
	}
	return strings.Join(out, "|")
}
