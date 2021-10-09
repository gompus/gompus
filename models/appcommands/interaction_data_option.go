package appcommands

// InteractionDataOption represents a command parameter.
// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-interaction-data-option-structure.
type InteractionDataOption struct {
	// Name is the parameter's name.
	Name string `json:"name,omitempty"`

	// Type is the parameter's type.
	Type OptionType `json:"type,omitempty"`

	// Value is the parameter's value.
	Value OptionType `json:"value,omitempty"`

	// Options is the set of options for groups and subcommands.
	Options []InteractionDataOption `json:"options,omitempty"`
}
