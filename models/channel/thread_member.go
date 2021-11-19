package channel

import (
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

// A ThreadMember is used to indicate whether a user has joined a thread.
// See https://discord.com/developers/docs/resources/channel#thread-member-object.
type ThreadMember struct {
	// Id is the thread's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// UserId is the user's id.
	UserID snowflake.Snowflake `json:"user_id,omitempty"`

	// JoinTimestamp represents the time the user last joined the thread.
	JoinTimestamp timestamp.Timestamp `json:"join_timestamp,omitempty"`

	// Flags contains user-thread settings.
	Flags int `json:"flags,omitempty"`
}
