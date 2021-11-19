package channel

import "github.com/gompus/snowflake"

// An Overwrite represents a single overwrite for a channel.
// See https://discord.com/developers/docs/resources/channel#overwrite-object.
type Overwrite struct {
	// Id is the id of the role or user.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Type is the overwrite's type.
	Type OverwriteType `json:"type,omitempty"`

	// Allow is the permission bit set to allow.
	Allow string `json:"allow,omitempty"`

	// Deny is the permission bit set to deny.
	Deny string `json:"deny,omitempty"`
}

//go:generate stringer -type=OverwriteType -trimprefix=OverWrite

// OverwriteType is used to distinguish between overwrites for roles and users.
type OverwriteType int

const (
	OverWriteRole   OverwriteType = 0
	OverWriteMember               = 1
)
