package commands

import (
	"encoding/json"
	"github.com/gompus/snowflake"
)

// RequestGuildMembers is the command for requesting
// all members for a guild or a list of guilds.
type RequestGuildMembers struct {
	// GuildID is the id of the guild for which to get members
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Query is the username prefix that returned members must have.
	Query string `json:"query,omitempty"`

	// Limit is the maximum number of members to send.
	Limit int `json:"limit,omitempty"`

	// Presences indicates whether the presences of the
	// matched members should be included in the response.
	Presences bool `json:"presences,omitempty"`

	// UserIDs can be used to specifiy which users to fetch.
	UserIDs UserIDs `json:"user_ids,omitempty"`

	// Nonce identifies the guild members chunk response.
	Nonce string `json:"nonce,omitempty"`
}

func CollectUserIDs(ids ...snowflake.Snowflake) UserIDs {
	return UserIDs{UserIDs: ids}
}

type UserIDs struct {
	UserIDs []snowflake.Snowflake
}

func (ids UserIDs) MarshalJSON() ([]byte, error) {
	if len(ids.UserIDs) == 1 {
		return json.Marshal(ids.UserIDs[0])
	}
	return json.Marshal(ids.UserIDs)
}
