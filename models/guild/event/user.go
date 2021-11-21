package event

import (
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
)

// User represents a user that has subscribed to an event.
// See https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-user-object.
type User struct {
	// EventID is the id of the event the user subcribed to.
	EventID snowflake.Snowflake `json:"guild_scheduled_event_id,omitempty"`

	// User is the user who subscribed to the event.
	User user.User `json:"user,omitempty"`

	// Member is the guild member representing the user.
	Member guild.Member `json:"member,omitempty"`
}
