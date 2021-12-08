package events

import (
	"fmt"
	"github.com/gompus/gompus/models/guild/event"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

type CreateParams struct {
	// ChannelID is the id of the event's channel.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// Metadata contains the event's entity metadata.
	Metadata event.EntityMetadata `json:"entity_metadata,omitempty"`

	// Name is the event's name.
	Name string `json:"name,omitempty"`

	// Privacy is the event's privacy level.
	PrivacyLevel event.PrivacyLevel `json:"privacy_level,omitempty"`

	// StartTime is the time at which the event starts.
	StartTime timestamp.Timestamp `json:"start_time,omitempty"`

	// EndTime is the time at which the event ends.
	EndTime timestamp.Timestamp `json:"end_time,omitempty"`

	// Description is the event's description.
	Description string `json:"description,omitempty"`

	// Type is the event's entity type.
	Type event.EntityType `json:"entity_type,omitempty"`
}

// Create creates a scheduled event in the guild with the given id.
func Create(token auth.Token, guildID snowflake.Snowflake) (event event.ScheduledEvent, err error) {
	return event, client.POST(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/scheduled-events", guildID),
		Result: &event,
		Token:  token,
	})
}
