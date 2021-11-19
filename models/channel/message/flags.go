package message

import "strings"

// Flags contain flags for a message.
type Flags int

const (
	CrosspostedFlag Flags = 1 << iota
	IsCrosspostFlag
	SuppressEmbedsFlag
	SourceMessageDeletedFlag
	UrgentFlag
	HasThreadFlag
	EphemeralFlag
	LoadingFlag
)

var flagsToString = map[Flags]string{
	CrosspostedFlag:          "crossposted",
	IsCrosspostFlag:          "is_crosspost",
	SuppressEmbedsFlag:       "suppress_embeds",
	SourceMessageDeletedFlag: "source_message_deleted",
	UrgentFlag:               "urgent",
	HasThreadFlag:            "has_thread",
	EphemeralFlag:            "ephemeral",
	LoadingFlag:              "loading",
}

// String returns the string representation of f.
func (f Flags) String() string {
	var flags []string
	for flag, name := range flagsToString {
		if f&flag != 0 {
			flags = append(flags, name)
		}
	}
	return strings.Join(flags, "|")
}
