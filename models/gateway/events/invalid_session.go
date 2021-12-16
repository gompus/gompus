package events

// InvalidSessionEvent is the event sent to
// indicate that a session has become invalid.
type InvalidSessionEvent struct {
	// Resumable indicates whether the session may be resumed.
	Resumable bool `json:"resumable,omitempty"`
}
