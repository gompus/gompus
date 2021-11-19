package mentions

import "github.com/gompus/snowflake"

// Allowed provides granular control over mentions.
// See https://discord.com/developers/docs/resources/channel#allowed-mentions-object.
type Allowed struct {
	// Parse is the set of allowed mention types to parse.
	Parse []Allowed `json:"parse,omitempty"`

	// Roles is the set of role ids to mention.
	Roles []snowflake.Snowflake `json:"roles,omitempty"`

	// Users is the set of user ids to mention.
	Users []snowflake.Snowflake `json:"users,omitempty"`

	// RepliedUser indicates whether to mention
	// the author of the message being replied to.
	RepliedUser bool `json:"replied_user,omitempty"`
}
