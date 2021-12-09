package guild

import (
	"fmt"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// GetMember retrieves the guild member for the user
// with the given id from the guild with the given id.
func GetMember(token auth.Token, guildID, userID snowflake.Snowflake) (member guild.Member, err error) {
	return member, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/members/%s", guildID, userID),
		Token:  token,
		Result: &member,
	})
}

type GetMembersParams struct {
	// Limit is the maximum number of members to retrieve.
	Limit int `json:"limit,omitempty"`

	// Ater is the maximum user id to include in the results.
	After snowflake.Snowflake `json:"after,omitempty"`
}

// GetMembers retrieves a set of guild
// members from the guild with the given id.
func GetMembers(token auth.Token, guildID snowflake.Snowflake, params ...GetMembersParams) (members []guild.Member, err error) {
	return members, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/members", guildID),
		Query:  client.GenerateQuery(params),
		Token:  token,
		Result: &members,
	})
}

type SearchMembersParams struct {
	// Limit is the maximum number of members to retrieve.
	Limit int `json:"limit,omitempty"`
}

// SearchMembers retrieves a set of guild members from the
// guild with the given id whose nicknames or usernames
// start with the given query string.
func SearchMembers(token auth.Token, query string, params ...SearchMembersParams) (members []guild.Member, err error) {
	return members, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/members/search", query),
		Query:  client.GenerateQuery(params),
		Token:  token,
		Result: &members,
	})
}
