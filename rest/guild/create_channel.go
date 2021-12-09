package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type CreateChannelParams struct {
	// Name is the channel's name.
	Name string `json:"name,omitempty"`

	// Type is the channel's type.
	Type channel.Type `json:"type,omitempty"`

	// Topic is the channel's topic.
	Topic string `json:"topic,omitempty"`

	// Bitrate is the channel's bitrate.
	Bitrate int `json:"bitrate,omitempty"`

	// UserLimit is the maximum amount of
	// users allowed in the voice channel.
	UserLimit int `json:"user_limit,omitempty"`

	// RateLimitPerUser is the amount of secons a user has
	// to wait before sending another message in the channel.
	RateLimitPerUser int `json:"rate_limit_per_user,omitempty"`

	// Position is the channel's position.
	Position int `json:"position,omitempty"`

	// PermissionOverwrites contains the channel's permission overwrites.
	PermissionOverwrites []channel.Overwrite `json:"permission_overwrites,omitempty"`

	// PrentID is the id of the channel's parent category.
	ParentID snowflake.Snowflake `json:"parent_id,omitempty"`

	// NSFW indicates whether the channel is nsfw.
	NSFW bool `json:"nsfw,omitempty"`
}

// CreateChannel creates a new channel
// in the guild with the given id.
func CreateChannel(token auth.Token, guildID snowflake.Snowflake, params CreateChannelParams) (channel channel.Channel, err error) {
	return channel, client.POST(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/channels", guildID),
		Body:   params,
		Token:  token,
		Result: &channel,
	})
}
