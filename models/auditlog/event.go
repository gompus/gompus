package auditlog

//go:generate stringer -type=Event

// Event represents an audit log event.
type Event int

const (
	GuildUpdateEvent            Event = 1
	ChannelCreateEvent                = 10
	ChannelUpdateEvent                = 11
	ChannelDeleteEvent                = 12
	ChannelOverwriteCreateEvent       = 13
	ChannelOverwriteUpdateEvent       = 14
	ChannelOverwriteDeleteEvent       = 15
	MemberKickEvent                   = 20
	MemberPruneEvent                  = 21
	MemberBanAddEvent                 = 22
	MemberBanRemoveEvent              = 23
	MemberUpdateEvent                 = 24
	MemberRoleUpdateEvent             = 25
	MemberMoveEvent                   = 26
	MemberDisconnectEvent             = 27
	BotAddEvent                       = 28
	RoleCreateEvent                   = 30
	RoleUpdateEvent                   = 31
	RoleDeleteEvent                   = 32
	InviteCreateEvent                 = 40
	InviteUpdateEvent                 = 41
	InviteDeleteEvent                 = 42
	WebhookCreateEvent                = 50
	WebhookUpdateEvent                = 51
	WebhookDeleteEvent                = 52
	EmojiCreateEvent                  = 60
	EmojiUpdateEvent                  = 61
	EmojiDeleteEvent                  = 62
	MessageDeleteEvent                = 72
	MessageBulkDeleteEvent            = 73
	MessagePinEvent                   = 74
	MessageUnpinEvent                 = 75
	IntegrationCreateEvent            = 80
	IntegrationUpdateEvent            = 81
	IntegrationDeleteEvent            = 82
	StageInstanceCreateEvent          = 83
	StageInstanceUpdateEvent          = 84
	StageInstanceDeleteEvent          = 85
	StickerCreateEvent                = 90
	StickerUpdateEvent                = 91
	StickerDeleteEvent                = 92
	ThreadCreateEvent                 = 110
	ThreadUpdateEvent                 = 111
	ThreadDeleteEvent                 = 112
)
