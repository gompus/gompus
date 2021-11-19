package message

// An Activity represents a message activity.
// See https://discord.com/developers/docs/resources/channel#message-object-message-activity-structure.
type Activity struct {
	// Type is the activitiy's type.
	Type Type `json:"type,omitempty"`

	// PartyID is the party id from a Rich Presence event.
	PartyID string `json:"party_id,omitempty"`
}
