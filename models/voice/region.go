package voice

// A Region represents a voice region.
// See https://discord.com/developers/docs/resources/voice#voice-region-object.
type Region struct {
	// Id is the region's id.
	ID string `json:"id,omitempty"`

	// Name is the region's name.
	Name string `json:"name,omitempty"`

	// Optimal indicates whether the region is optimal for the client.
	Optimal bool `json:"optimal,omitempty"`

	// Deprecated indicates whether the region is deprecated.
	Deprecated bool `json:"deprecated,omitempty"`

	// Custom indicates whether the region is custom.
	Custom bool `json:"custom,omitempty"`
}
