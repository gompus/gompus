package events

import (
	"github.com/gompus/gompus/models/gateway"
	"github.com/gompus/gompus/models/gateway/commands"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
)

// PresenceUpdateEvent is the event sent when a user's
// presence or info, such as their name or avatar, is
// updated.
type PresenceUpdateEvent struct {
	// User is the user whose presence is being updated.
	User user.User `json:"user,omitempty"`

	// GuildID is the of the user's guild.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Status is the user's new status.
	Status commands.Status `json:"status,omitempty"`

	// Activites contains the user's current activities.
	Activites []gateway.Activity `json:"activites,omitempty"`

	// ClientStatus is the user's platform-dependent status.
	ClientStatus ClientStatus `json:"client_status,omitempty"`
}

// ClientStatus represents a user's platform-dependent status.
type ClientStatus struct {
	// Desktop is the user's status set for
	// an active desktop application session.
	Desktop string `json:"desktop,omitempty"`

	// Mobile is the user's status set for
	// an active mobile application session.
	Mobile string `json:"mobile,omitempty"`

	// Web is the user's status set for
	// an active web application session.
	Web string `json:"web,omitempty"`
}
