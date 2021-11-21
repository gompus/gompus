package event

//go:generate stringer -type=EntityType -trimprefix=EntityType

// EntityType represents the type of an entity affected by an event.
type EntityType int

const (
	EntityTypeStageInstance EntityType = iota + 1
	EntityTypeVoice
	EntityTypeExternal
)
