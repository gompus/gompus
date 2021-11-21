package event

//go:generate stringer -type=Status -trimprefix=Status

// Status represents the status of an event.
type Status int

const (
	StatusScheduled Status = iota + 1
	StatusActive
	StatusCompleted
	StatusCanceled
)
