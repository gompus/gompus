package channel

//go:generate stringer -type=Type -trimprefix=Type

// Type represents the type of a Channel.
type Type int

const (
	// TypeGuildText is the type for text channels.
	TypeGuildText Type = 0

	// TypeDM is the type for direct messages.
	TypeDM = 1

	// TypeGuildVoice is the type for voice channels.
	TypeGuildVoice = 2

	// TypeGroupDM is the type for direct
	// messages between multiple users.
	TypeGroupDM = 3

	// TypeGuildCategory is the type for organizational
	// categories that contain up to 50 channels.
	TypeGuildCategory = 4

	// TypeGuildNews is the type for channels that users
	// can follow and crosspost into their own server.
	TypeGuildNews = 5

	// TypeGuildStore is the type for channels in
	// which game developers can sell their games.
	TypeGuildStore = 6

	// TypeGuildNewsThread is the type for temporary
	// sub-channels within guild news channels.
	TypeGuildNewsThread = 10

	// TypeGuildPublicThread is the type for temporary
	// sub-channels withing guild text channels.
	TypeGuildPublicThread = 11

	// TypeGuildPrivateThread is the type for temporary
	// sub-channels within guild text channels that are
	// only viewable by those invited and those with the
	// manage threads permission.
	TypeGuildPrivateThread = 12

	// TypeGuildStageVoice is the type for voice channels
	// used for hosting events with an audience.
	TypeGuildStageVoice = 13
)
