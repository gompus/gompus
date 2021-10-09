package appcommands

// OptionType is the type of an Option.
type OptionType int

const (
	SubCommandOptions OptionType = iota + 1
	SubCommandGroupOptions
	StringOptions
	IntegerOptions
	BooleanOptions
	UserOptions
	ChannelOptions
	RoleOptions
	MentionableOptions
	NumberOptions
)
