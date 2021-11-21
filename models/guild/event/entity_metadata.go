package event

// EntityMetadata contains metadata about an event.
// See https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-entity-metadata.
type EntityMetadata struct {
	// Location represents the event's location.
	Location string `json:"location,omitempty"`
}
