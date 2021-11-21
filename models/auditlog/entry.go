package auditlog

import "github.com/gompus/snowflake"

// Entry represents an audit log entry.
// See https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object.
type Entry struct {
	// TargetID is the id of the affected entity.
	TargetID *string `json:"target_id,omitempty"`

	// Changes contains the changes made to the affected entity.
	Changes []Change `json:"changes,omitempty"`

	// UserID is the id of the user who made the changes.
	UserID *snowflake.Snowflake `json:"user_id,omitempty"`

	// ID is the entry's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// ActionType is the type of action that occurred.
	ActionType Event `json:"action_type,omitempty"`

	// Options contains additional info for certain action types.
	Options OptionalInfo `json:"options,omitempty"`

	// Reason is the reason for the change.
	Reason string `json:"reason,omitempty"`
}
