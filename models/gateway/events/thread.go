package events

import (
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/snowflake"
)

// ThreadCreateEvent is the event send when a thread is
// created, or when the current user is added to a thread.
type ThreadCreateEvent struct {
	channel.Channel
	channel.ThreadMember
}

// ThreadDeleteEvent is the event sent when a thread is deleted.
type ThreadDeleteEvent struct {
	// ID is the channel's id.
	ChannelID snowflake.Snowflake `json:"id,omitempty"`

	// GuildID is the id of the guild the channel belonged to.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// ParentID is the id of the channel's parent
	// category (only valid for guild channels).
	// For threads, ParentID is the id of the text
	// channel in which the thread was created.
	ParentID snowflake.Snowflake `json:"parent_id,omitempty"`

	// Type is the channel's type.
	Type channel.Type `json:"type,omitempty"`
}

// ThreadListSyncEvent is the event sent when
// the current user gains access to a channel.
type ThreadListSyncEvent struct {
	// GuildID is the id of the channel's guild.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// ChannelIDs contains the parent channel id's whose threads
	// are being synced. If omitted, threads were synced for the
	// entire guild.
	ChannelIDs []snowflake.Snowflake `json:"channel_ids,omitempty"`

	// Threads contains all active threads in the
	// given channels that the current user can access.
	Threads []channel.Channel `json:"threads,omitempty"`

	// Members contains all thread members from the synced
	// threads for the current user, indicating which threads
	// the current user has been added to.
	Members []channel.ThreadMember `json:"members,omitempty"`
}

// ThreadMemberUpdateEvent is the event sent when
// the thread member for the current user is updated.
type ThreadMemberUpdateEvent channel.ThreadMember

// ThreadMembersUpdateEvent is the event sent when
// anyone is added to or removed from a thread.
type ThreadMembersUpdateEvent struct {
	// ID is the thread's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// GuildID is the id of the thread's guild.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// MemberCount is the approximate number of members in the
	// thread. Note that no more than 50 members are counted.
	MemberCount int `json:"member_count,omitempty"`

	// AddedMembers contains the users that were added to the thread.
	AddedMembers []channel.ThreadMember `json:"added_members,omitempty"`

	// RemovedMemberIDs contains the id of the
	// users who were removed from the thread.
	RemovedMemberIDs []snowflake.Snowflake `json:"removed_member_ids,omitempty"`
}

// ThreadUpdateEvent is the event sent when a thread is updated.
type ThreadUpdateEvent channel.Channel
