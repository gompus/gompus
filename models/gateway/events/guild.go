package events

import (
	"github.com/gompus/gompus/models/emoji"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/models/guild/event"
	"github.com/gompus/gompus/models/permissions"
	"github.com/gompus/gompus/models/sticker"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

// GuildCreateEvent is the even sent when...
//
//  1. a user is initially connecting to the gateway server.
//  2. a guild becomes available to the client.
//  3. the current user joins a new guild.
//
type GuildCreateEvent guild.Guild

// GuildDeleteEvent is the event sent when a guild becomes
// or was already unavailable due to an outage, or when the
// user leaves or is removed from a guild.
type GuildDeleteEvent UnavailableGuild

// GuildUpdateEvent is the event sent when a guild is updated.
type GuildUpdateEvent guild.Guild

// GuildBanAddEvent is the event sent when a user is banned from a guild.
type GuildBanAddEvent struct {
	// GuildID is the id of the guild the user was banned from.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// User is the banned user.
	User user.User `json:"user,omitempty"`
}

// GuildBanRemoveEvent is the event sent when a user is unbanned from a guild.
type GuildBanRemoveEvent struct {
	// GuildID is the id of the guild the user was unbanned from.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// User is the unbanned user.
	User user.User `json:"user,omitempty"`
}

// GuildEmojisUpdateEvent is the event sent
// when a guild's emojis have been updated.
type GuildEmojisUpdateEvent struct {
	// GuildID is the id of the guild whose emojis have been updated.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Emojis contains the updated emojis.
	Emojis []emoji.Emoji `json:"emojis,omitempty"`
}

// GuildStickersUpdateEvent is the event sent
// when a guild's stickers have been updated.
type GuildStickersUpdateEvent struct {
	// GuildID is the id of the guild whose stickers have been updated.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Stickers contains the updated stickers.
	Stickers []sticker.Sticker `json:"stickers,omitempty"`
}

// GuildIntegrationsUpdateEvent is the event
// sent when a guild integration is updated.
type GuildIntegrationsUpdateEvent struct {
	// GuildID is the id of the guild whose integration has been updated.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`
}

// GuildMemberAddEvent is the event sent when a new user joins a guild.
type GuildMemberAddEvent struct {
	guild.Member

	// GuildID is the id of the guild the user joined.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`
}

// GuildMemberRemoveEvent is the event sent when a
// user is removed from a guild (leave, kick, ban).
type GuildMemberRemoveEvent struct {
	// GuildID is the id of the guild the user was removed from.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// User is the removed user.
	User user.User `json:"user,omitempty"`
}

// GuildMemberUpdateEvent is the event
// sent when a guild member is updated.
type GuildMemberUpdateEvent struct {
	// GuildID is the id of the guild whose member was updated.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Roles contains user role ids.
	Roles []snowflake.Snowflake `json:"roles,omitempty"`

	// User is the updated user.
	User user.User `json:"user,omitempty"`

	// Nick is the user's nick name in the guild.
	Nick *string `json:"nick,omitempty"`

	// Avatar is the user's guild avatar hash.
	Avatar *string `json:"avatar,omitempty"`

	// JoinedAt is the time at which the user joined the guild.
	JoinedAt *timestamp.Timestamp `json:"joined_at,omitempty"`

	// PremiumSince is the guild at which
	// the user started boosting the guild.
	PremiumSince *timestamp.Timestamp `json:"premium_since,omitempty"`

	// Deaf indicates whether the user is deafened in voice channels.
	Deaf bool `json:"deaf,omitempty"`

	// Mute indicates whether the useris muted in voice channels.
	Mute bool `json:"mute,omitempty"`

	// Pending indicates whether the user has not yet
	// passed the guild's membership sreening requirements.
	Pending bool `json:"pending,omitempty"`
}

// GuildMembersChunkEvent is the event sent in
// response to reques guild members commands.
type GuildMembersChunkEvent struct {
	// GuildID is the guild's id.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Members contains the guild's members.
	Members []guild.Member `json:"members,omitempty"`

	// ChunkIndex is the chunk index in the expected chunks.
	ChunkIndex int `json:"chunk_index,omitempty"`

	// ChunkCount is the total number of expected chunks.
	ChunkCount int `json:"chunk_count,omitempty"`

	// NotFound contains invalid ids passed
	// to the request guild members command.
	NotFound []snowflake.Snowflake `json:"not_found,omitempty"`

	// Presences contains the presences of the returned members.
	Presences []PresenceUpdateEvent `json:"presences,omitempty"`

	// Nonce is the nonce used in the guild members request.
	Nonce string `json:"nonce,omitempty"`
}

// GuildRoleCreateEvent is the event sent when a guild role is created.
type GuildRoleCreateEvent struct {
	// GuildID is the id of the guild the role was created in.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Role is the created role.
	Role permissions.Role `json:"role,omitempty"`
}

// GuildRoleUpdateEvent is the event sent when a guild role is updated.
type GuildRoleUpdateEvent struct {
	// GuildID is the id of the guild the role was updated in.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Role is the updated role.
	Role permissions.Role `json:"role,omitempty"`
}

// GuildRoleDeleteEvent is the event sent when a guild role is deleted.
type GuildRoleDeleteEvent struct {
	// GuildID is the id of the guild the role was deleted in.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// RoleID is the id of the deleted role.
	RoleID snowflake.Snowflake `json:"role_id,omitempty"`
}

// GuildScheduledEventCreateEvent is the event
// sent when a guild scheduled event is created.
type GuildScheduledEventCreateEvent event.ScheduledEvent

// GuildScheduledEventUpdateEvent is the event
// sent when a guild scheduled event is updated.
type GuildScheduledEventUpdateEvent event.ScheduledEvent

// GuildScheduledEventDeleteEvent is the event
// sent when a guild scheduled event is deleted.
type GuildScheduledEventDeleteEvent event.ScheduledEvent

// GuildScheduledEventUserAddEvent is the event sent
// when a user has subscribed to a guild scheduled event.
type GuildScheduledEventUserAddEvent struct {
	// EventID is the scheduled event's id.
	EventID snowflake.Snowflake `json:"guild_scheduled_event_id,omitempty"`

	// UserID is the id of the user that subscribed to the event.
	UserID snowflake.Snowflake `json:"user_id,omitempty"`

	// GuildID is the id of the guild scheduling the event.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`
}

// GuildScheduledEventUserRemoveEvent is the event sent when
// a user has unsubscribed from a guild scheduled event.
type GuildScheduledEventUserRemoveEvent struct {
	// EventID is the scheduled event's id.
	EventID snowflake.Snowflake `json:"guild_scheduled_event_id,omitempty"`

	// UserID is the id of the user that unsubscribed to the event.
	UserID snowflake.Snowflake `json:"user_id,omitempty"`

	// GuildID is the id of the guild scheduling the event.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`
}
