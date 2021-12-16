package events

import (
	"github.com/gompus/gompus/models/application"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
)

// ReadyEvent is the event sent when a client has completed
// the initial handshake with the gateway. It contains all
// the state required for a client to begin interacting with
// the rest of the platform.
type ReadyEvent struct {
	Version    int                `json:"v,omitempty"`
	User       user.User          `json:"user,omitempty"`
	Guilds     []UnavailableGuild `json:"guilds,omitempty"`
	SessionID  string             `json:"session_id,omitempty"`
	Shard      [2]int             `json:"shard,omitempty"`
	Appliction PartialApplication `json:"appliction,omitempty"`
}

type UnavailableGuild struct {
	ID          snowflake.Snowflake `json:"id,omitempty"`
	Unavailable bool                `json:"unavailable,omitempty"`
}

type PartialApplication struct {
	ID    snowflake.Snowflake `json:"id,omitempty"`
	Flags application.Flags   `json:"flags,omitempty"`
}
