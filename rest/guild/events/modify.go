package events

import (
	"fmt"
	"github.com/gompus/gompus/models/guild/event"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

type ModifyParams struct {
	// ChannelID is the id of the event's channel.
	ChannelID *snowflake.Snowflake `json:"channel_id,omitempty"`

	// Metadata contains the event's metadata.
	Metadata event.EntityMetadata `json:"metadata,omitempty"`

	// Name is the event's name.
	Name string `json:"name,omitempty"`

	// PrivacyLevel is the event's privay level.
	PrivacyLevel event.PrivacyLevel `json:"privacy_level,omitempty"`

	// StartTime is the time the event starts.
	StartTime timestamp.Timestamp `json:"scheduled_start_time,omitempty"`

	// EndTime is the time the event ends.
	EndTime timestamp.Timestamp `json:"scheduled_end_time,omitempty"`

	// Description is the event's description.
	Description string `json:"description,omitempty"`

	// Type is the event's entity type.
	Type event.EntityType `json:"entity_type,omitempty"`

	// Status is the event's status.
	Status event.Status `json:"status,omitempty"`
}

// Modify modifies the event with the given
// id in the guild with the given id.
func Modify(token auth.Token, guildID, eventID snowflake.Snowflake, params ModifyParams) (event event.ScheduledEvent, err error) {
	return event, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/scheduled-events/%s", guildID, eventID),
		Body:   params,
		Token:  token,
		Result: &event,
	})
}
