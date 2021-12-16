package gateway

import "strings"

// Intents allow tighter control over what types
// of gateway intents a user can subscribe to.
// See https://discord.com/developers/docs/topics/gateway#gateway-intents.
type Intents int

const (
	GuildsIntents Intents = 1 << iota
	GuildMembersIntents
	GuildBansIntents
	GuildEmojisAndStickersIntents
	GuildIntegrationsIntents
	GuildWebhooksIntents
	GuildInvitesIntents
	GuildVoiceStatesIntents
	GuildPresencesIntents
	GuildMessagesIntents
	GuildMessageReactionsIntents
	GuildMessageTypingIntents
	DirectMessagesIntents
	DirectMessageReactionsIntents
	DirectMessageTypingIntents
	_
	GuildScheduledEventsIntents
)

var intentToStr = map[Intents]string{
	GuildsIntents:                 "Guilds",
	GuildMembersIntents:           "GuildMembers",
	GuildBansIntents:              "GuildBans",
	GuildEmojisAndStickersIntents: "GuildEmojisAndStickers",
	GuildIntegrationsIntents:      "GuildIntegrations",
	GuildWebhooksIntents:          "GuildWebhooks",
	GuildInvitesIntents:           "GuildInvites",
	GuildVoiceStatesIntents:       "GuildVoiceStates",
	GuildPresencesIntents:         "GuildPresences",
	GuildMessagesIntents:          "GuildMessages",
	GuildMessageReactionsIntents:  "GuildMessageReactions",
	GuildMessageTypingIntents:     "GuildMessageTyping",
	DirectMessagesIntents:         "DirectMessages",
	DirectMessageReactionsIntents: "DirectMessageReactions",
	DirectMessageTypingIntents:    "DirectMessageTyping",
	GuildScheduledEventsIntents:   "GuildScheduledEvents",
}

// String returns the string representation of i.
func (i Intents) String() string {
	var out []string
	for flag, s := range intentToStr {
		if i&flag != 0 {
			out = append(out, s)
		}
	}
	return strings.Join(out, "|")
}
