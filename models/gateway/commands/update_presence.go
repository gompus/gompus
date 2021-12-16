package commands

import "github.com/gompus/gompus/models/gateway"

// UpdatePresence is the command for indicating a presence or status update.
// See https://discord.com/developers/docs/topics/gateway#update-presence.
type UpdatePresence struct {
	// Since is the unix time (in ms) of then the client went idle.
	Since *int `json:"since,omitempty"`

	// Activities contains the user's activities.
	Activities []gateway.Activity `json:"activities,omitempty"`

	// Status is the user's new status.
	Status Status `json:"status,omitempty"`

	// AFK indicates whetehr the client is afk.
	AFK bool `json:"afk,omitempty"`
}

// Status represents a user's status.
type Status string

const (
	StatusOnline       Status = "online"
	StatusIdle         Status = "idle"
	StatusDoNotDisturb Status = "dnd"
	StatusInvisible    Status = "invisible"
	StatusOffline      Status = "offline"
)
