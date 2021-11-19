package embed

// Field represents a field of an embed.
// See https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure.
type Field struct {
	// Name is the field's name.
	Name string `json:"name,omitempty"`

	// Value is the field's value.
	Value string `json:"value,omitempty"`

	// Inline indicates whether the field should be displayed inline.
	Inline bool `json:"inline,omitempty"`
}
