package application

//go:generate stringer -type=Flag

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
