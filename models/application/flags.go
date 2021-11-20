package application

import "strings"

// Flags describe an Application.
type Flags int

const (
	FlagsGatewayPresence               Flags = 1 << 12
	FlagsGatewayPresenceLimited              = 1 << 13
	FlagsGatewayGuildMembers                 = 1 << 14
	FlagsGatewayGuildMembersLimited          = 1 << 15
	FlagsVerificationPendingGuildLimit       = 1 << 16
	FlagsEmbedded                            = 1 << 17
	FlagsGatewayMessageContent               = 1 << 18
	FlagsGatewayMessageContentLimited        = 1 << 19
)

var flagToStr = map[Flags]string{
	FlagsGatewayPresence:               "GatewayPresence",
	FlagsGatewayPresenceLimited:        "GatewayPresenceLimited",
	FlagsGatewayGuildMembers:           "GatewayGuildMembers",
	FlagsGatewayGuildMembersLimited:    "GatewayGuildMembersLimited",
	FlagsVerificationPendingGuildLimit: "VerificationPendingGuildLimit",
	FlagsEmbedded:                      "Embedded",
	FlagsGatewayMessageContent:         "GatewayMessageContent",
	FlagsGatewayMessageContentLimited:  "GatewayMessageContentLimited",
}

// String returns a string representation of f.
func (f Flags) String() string {
	var out []string
	for flag, s := range flagToStr {
		if f&flag != 0 {
			out = append(out, s)
		}
	}
	return strings.Join(out, "|")
}
