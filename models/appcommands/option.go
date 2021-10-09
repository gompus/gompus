package appcommands

// Option represents a parameter for an ApplicationCommand.
type Option struct {
	// Type is the option's Type.
	Type OptionType `json:"type,omitempty"`

	// Name is the option's name.
	Name string `json:"name,omitempty"`

	// Description is the option's description.
	Description string `json:"description,omitempty"`

	// Required reports whether the option is required.
	Required bool `json:"required,omitempty"`

	// Choices is the set of valid choices the user can select from.
	Choices []OptionChoiceValue `json:"choices,omitempty"`

	// Options is the set of parameters for
	// subcommand or subcommand group options.
	Options []Option `json:"options,omitempty"`

	// ChannelTypes is the set of channel types to
	// show if the option is of type ChannelOptions.
	ChannelTypes []interface{} `json:"channel_types,omitempty"`
}
