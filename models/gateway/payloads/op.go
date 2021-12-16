package payloads

//go:generate stringer -type=Op -trimprefix=Op

// Op represents a payload's opcode.
type Op int

const (
	OpDispatch            Op = 0
	OpHeartbeat              = 1
	OpIdentify               = 2
	OpPresenceUpdate         = 3
	OpVoiceStateUpdate       = 4
	OpResume                 = 6
	OpReconnect              = 7
	OpRequestGuildMembers    = 8
	OpInvalidSession         = 9
	OpHelloReceive           = 10
	OpHeartbeatACK           = 11
)
