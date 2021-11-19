package message

//go:generate stringer -type=Type -trimprefix=Type

// Type is the type of a Message.
type Type int

const (
	TypeDefault Type = iota
	TypeRecipientAdd
	TypeRecipientRemove
	TypeCall
	TypeChannelNameChange
	TypeChannelIconChange
	TypeChannelPinnedMessage
	TypeGuildMemberJoin
	TypeUserPremiumGuildSubscription
	TypeUserPremiumGuildSubscriptionTier1
	TypeUserPremiumGuildSubscriptionTier2
	TypeUserPremiumGuildSubscriptionTier3
	TypeChannelFollowAdd
	_
	TypeGuildDiscoveryDisqualified
	TypeGuildDiscoveryRequalified
	TypeGuildDiscoveryGracePeriodInitialWarning
	TypeGuildDiscoveryGracePeriodFinalWarning
	TypeThreadCreated
	TypeReply
	TypeChatInputCommand
	TypeThreadStarterMessage
	TypeGuildInviteReminder
	TypeContextMenuCommand
)
