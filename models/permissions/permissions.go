package permissions

import "strings"

// Permissions represent a Discord user's permissions.
// See https://discord.com/developers/docs/topics/permissions#permissions.
type Permissions int64

const (
	CreateInstantInvite Permissions = 1 << iota
	KickMembers
	BanMembers
	Administrator
	ManageChannels
	ManageGuild
	AddReactions
	ViewAuditLog
	PrioritySpeaker
	Stream
	ViewChannel
	SendMessages
	SendTtsMessages
	ManageMessages
	EmbedLinks
	AttachFiles
	ReadMessageHistory
	MentionEveryone
	UseExternalEmojis
	ViewGuildInsights
	Connect
	Speak
	MuteMembers
	DeafenMembers
	MoveMembers
	UseVad
	ChangeNickname
	ManageNicknames
	ManageRoles
	ManageWebhooks
	ManageEmojisAndStickers
	UseApplicationCommands
	RequestToSpeak
	ManageThreads
	CreatePublicThreads
	CreatePrivateThreads
	UseExternalStickers
	SendMessagesInThreads
	StartEmbeddedActivities
)

var permissionToString = map[Permissions]string{
	CreateInstantInvite:     "CreateInstantInvite",
	KickMembers:             "KickMembers",
	BanMembers:              "BanMembers",
	Administrator:           "Administrator",
	ManageChannels:          "ManageChannels",
	ManageGuild:             "ManageGuild",
	AddReactions:            "AddReactions",
	ViewAuditLog:            "ViewAuditLog",
	PrioritySpeaker:         "PrioritySpeaker",
	Stream:                  "Stream",
	ViewChannel:             "ViewChannel",
	SendMessages:            "SendMessages",
	SendTtsMessages:         "SendTtsMessages",
	ManageMessages:          "ManageMessages",
	EmbedLinks:              "EmbedLinks",
	AttachFiles:             "AttachFiles",
	ReadMessageHistory:      "ReadMessageHistory",
	MentionEveryone:         "MentionEveryone",
	UseExternalEmojis:       "UseExternalEmojis",
	ViewGuildInsights:       "ViewGuildInsights",
	Connect:                 "Connect",
	Speak:                   "Speak",
	MuteMembers:             "MuteMembers",
	DeafenMembers:           "DeafenMembers",
	MoveMembers:             "MoveMembers",
	UseVad:                  "UseVad",
	ChangeNickname:          "ChangeNickname",
	ManageNicknames:         "ManageNicknames",
	ManageRoles:             "ManageRoles",
	ManageWebhooks:          "ManageWebhooks",
	ManageEmojisAndStickers: "ManageEmojisAndStickers",
	UseApplicationCommands:  "UseApplicationCommands",
	RequestToSpeak:          "RequestToSpeak",
	ManageThreads:           "ManageThreads",
	CreatePublicThreads:     "CreatePublicThreads",
	CreatePrivateThreads:    "CreatePrivateThreads",
	UseExternalStickers:     "UseExternalStickers",
	SendMessagesInThreads:   "SendMessagesInThreads",
	StartEmbeddedActivities: "StartEmbeddedActivities",
}

// String returns the string representation of p.
func (p Permissions) String() string {
	var out []string
	for perm, s := range permissionToString {
        if p&perm != 0 {
            out = append(out, s)
        }
    }
	return strings.Join(out, "|")
}
