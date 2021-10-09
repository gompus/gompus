package appcommands

// Type represents the type of an ApplicationCommand.
// See https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types.
type Type int

const (
	// TypeChatInput is the Type for slash commands.
	TypeChatInput Type = iota + 1

	// TypeUser is the Type for UI-based commands that show
	// up when the user right clicks or taps on another user.
	TypeUser

	// TypeMessage is the Type for UI-based commands that show
	// up when the user right clicks or taps on a message.
	TypeMessage
)
