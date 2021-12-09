package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type ModifyMemberParams struct {
	// Nick is the nickname to assign to the member.
	Nick string `json:"nick,omitempty"`

	// Roles contains the roles to assign the member.
	Roles []snowflake.Snowflake `json:"roles,omitempty"`

	// ChannelID is the id of the
	// channel to move the user to.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// Mute indicates whether to mute
	// the user in voice channels.
	Mute bool `json:"mute,omitempty"`

	// Deaf indicates whether to deafen
	// the user in voice channels.
	Deaf bool `json:"deaf,omitempty"`
}

// ModifyMember modifies the member with the
// given id in the guild with the given id.
func ModifyMember(token auth.Token, guildID, userID snowflake.Snowflake, params ModifyMemberParams) (member guild.Member, err error) {
	return member, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/members/%s", guildID, userID),
		Body:   params,
		Token:  token,
		Result: &member,
	})
}

type ModifyCurrentMemberParams struct {
	// Nick is the nickname to assign to the member.
	Nick *string `json:"nick,omitempty"`
}

// ModifyCurrentMember modifies the current
// member in the guild with the given id.
func ModifyCurrentMember(token auth.Token, guildID snowflake.Snowflake, params ModifyCurrentMemberParams) (member guild.Member, err error) {
	return member, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/members/@me", guildID),
		Body:   params,
		Token:  token,
		Result: &member,
	})
}
