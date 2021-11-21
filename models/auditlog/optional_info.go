package auditlog

import "github.com/gompus/snowflake"

// OptionalInfo contains optional information about an audit log entry.
// See https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-optional-audit-entry-info.
type OptionalInfo struct {
	// ChannelID is the id of the channel in which entities were targeted.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// Count is the number of entities that were targeted.
	Count string `json:"count,omitempty"`

	// DelteMemberDays is the number of days
	// after which inactive members were kicked.
	DeleteMemberDays string `json:"delete_member_days,omitempty"`

	// ID is the id of the overwriten entity.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// MemebrsRemoved is the number of members removed by the prune.
	MembersRemoved string `json:"members_removed,omitempty"`

	// MessageID is the id of the deleted message.
	MessageID snowflake.Snowflake `json:"message_id,omitempty"`

	// RoleName is the name of the affected role.
	RoleName string `json:"role_name,omitempty"`

	// Type is the type of the overwriten entity.
	Type string `json:"type,omitempty"`
}
