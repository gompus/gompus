package integration

import (
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
)

// Integration represents a Discord integration.
// See https://discord.com/developers/docs/resources/guild#integration-object.
type Integration struct {
	// ID is the integration's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Name is the integration's name.
	Name string `json:"name,omitempty"`

	// Tye is the integration's type.
	Type Type `json:"type,omitempty"`

	// Enabled indicaes whether the integration is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// Syncing reports whether the integration is syncing.
	Syncing bool `json:"syncing,omitempty"`

	// RoleID is the id the integration uses for "subscribers".
	RoleID snowflake.Snowflake `json:"role_id,omitempty"`

	// EnableEmoticons indicates whether emotes should be
	// synced for the integration (only for twitch integrations).
	EnableEmoticons bool `json:"enable_emoticons,omitempty"`

	// ExpireBehavior is the behaviour for expiring subscribers.
	ExpireBehavior ExpireBehavior `json:"expire_behavior,omitempty"`

	// ExpireGracePeriod is the grace period
	// (in days) before expiring subscribers.
	ExpireGracePeriod int `json:"expire_grace_period,omitempty"`

	// User is the integration's user.
	User user.User `json:"user,omitempty"`

	// Account contains the integration's account information.
	Account Account `json:"account,omitempty"`

	// SyncedAt is the time the integration was last synced.
	SyncedAt timestamp.Timestamp `json:"synced_at,omitempty"`

	// SubscriberCount is the number of subscribers of the integration.
	SubscriberCount int `json:"subscriber_count,omitempty"`

	// Revoked indicates whether the integration has been revoked.
	Revoked bool `json:"revoked,omitempty"`

	// Application is the bot/OAuth2 application for integrations.
	Application Application `json:"application,omitempty"`
}
