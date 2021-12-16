package commands

// Resume is the command for replaying missed
// events when a disconnected client resumes.
// See https://discord.com/developers/docs/topics/gateway#resume.
type Resume struct {
	// Token is the client's session token.
	Token string `json:"token,omitempty"`

	// SessionID is the client's session id.
	SessionId string `json:"session_id,omitempty"`

	// SeqNum is the last sequence number the client received.
	SeqNum int `json:"seq,omitempty"`
}
