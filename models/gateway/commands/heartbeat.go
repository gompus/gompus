package commands

// Heartbeat is the command for maintaining an active gateway connection.
// See https://discord.com/developers/docs/topics/gateway#heartbeat.
type Heartbeat struct {
	// SeqNum is the last sequence number the client received.
	SeqNum int `json:"d,omitempty"`
}
