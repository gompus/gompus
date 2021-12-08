package thread

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type StartParams struct {
	// Name is the channel's name.
	Name string `json:"name,omitempty"`

	// AutoArchiveDuration is the duration (in minutes)
	// after which to automatically archive the thread
	// after recent activity.
	AutoArchiveDuration int `json:"auto_archive_duration,omitempty"`

	// Type is type of thread to create.
	Type channel.Type `json:"type,omitempty"`

	// Invitable indicates whether non-moderators
	// can add other non-moderators to the thread.
	Invitable bool `json:"invitable,omitempty"`

	// RateLimitPerUser is the amount of seconds
	// a user has to wait before sending another
	// message.
	RateLimitPerUser *int `json:"rate_limit_per_user,omitempty"`
}

// Start creates a new thread in the channel with the
// given id that is not connected to an existing message.
func Start(token auth.Token, channelID snowflake.Snowflake, params StartParams) (channel channel.Channel, err error) {
	return channel, client.POST(client.Request{
		Path:   fmt.Sprintf("/channels/%s/threads", channelID),
		Body:   params,
		Token:  token,
		Result: &channel,
	})
}

type StartWithMessageParams struct {
	// Name is the channel's name.
	Name string `json:"name,omitempty"`

	// AutoArchiveDuration is the duration (in minutes)
	// after which to automatically archive the thread
	// after recent activity.
	AutoArchiveDuration int `json:"auto_archive_duration,omitempty"`

	// RateLimitPerUser is the amount of seconds
	// a user has to wait before sending another
	// message.
	RateLimitPerUser *int `json:"rate_limit_per_user,omitempty"`
}

// StartWithMessage creates a new thread in the channel
// with the given id from the message with the given id.
func StartWithMessage(token auth.Token, channelID, messageID snowflake.Snowflake, params StartWithMessageParams) (channel channel.Channel, err error) {
	return channel, client.POST(client.Request{
		Path:   fmt.Sprintf("/channels/%s/messages/%s/threads", channelID, messageID),
		Body:   params,
		Token:  token,
		Result: &channel,
	})
}
