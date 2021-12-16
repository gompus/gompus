package events

// HelloEvent is the event set upon connecting to a websocket server.
type HelloEvent struct {
	// HeartbeatInterval is the interval (in ms)
	// the client should heartbeat with.
	HeartbeatInterval int `json:"heartbeat_interval,omitempty"`
}
