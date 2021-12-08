package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// FollowNewsChannel follows the news channel with the given
// id to send messages to the target channel with the given id.
func FollowNewsChannel(token auth.Token, channelID, targetChannelID snowflake.Snowflake) (channel channel.FollowedChannel, err error) {
	return channel, client.POST(client.Request{
		Path: fmt.Sprintf("/channel/%s/follow", channelID),
		Body: struct {
			TargetChannelID snowflake.Snowflake `json:"target_channel_id"`
		}{
			TargetChannelID: targetChannelID,
		},
		Token:  token,
		Result: &channel,
	})
}
