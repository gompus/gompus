package mentions

import (
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/snowflake"
)

// A Mention represents a mention in a channel.
// See https://discord.com/developers/docs/resources/channel#channel-mention-object.
type Mention struct {
	// ID is the channel's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// GuildID is the id of the channel's guild.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Type is the channel's type.
	Type channel.Type `json:"type,omitempty"`

	// Name is the channel's name.
	Name string `json:"name,omitempty"`
}
