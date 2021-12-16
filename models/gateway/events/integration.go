package events

import (
	"github.com/gompus/gompus/models/guild/integration"
	"github.com/gompus/snowflake"
)

// IntegrationCreateEvent is the event sent when an integration is created.
type IntegrationCreateEvent struct {
	integration.Integration

	// GuildID is the id of the guild for which the integration was created.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`
}

// IntegrationUpdateEvent is the event sent when an integration is updated.
type IntegrationUpdateEvent struct {
	integration.Integration

	// GuildID is the id of the guild for which the integration was updated.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`
}

// IntegrationDeleteEvent is the event sent when an integration is deleted.
type IntegrationDeleteEvent struct {
	// ID is the id of the deleted integration.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// GuildID is the id of the guild for which the integration was deleted.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// ApplicationID is the id of the bot/OAuth2
	// application associated with the integration.
	ApplicationID snowflake.Snowflake `json:"application_id,omitempty"`
}
