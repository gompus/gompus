package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/invite"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type CreateInviteParams struct {
	// MaxAge is the time (in seconds)
	// after which the invite expires.
	MaxAge int `json:"max_age,omitempty"`

	// MaxUses is the maximum number
	// of uses (0 for unlimited).
	MaxUses int `json:"max_uses,omitempty"`

	// Temporary indicates whether the invite
	// only grants temporary membership.
	Temporary bool `json:"temporary,omitempty"`

	// Unique indicates whether the invite is unique.
	Unique bool `json:"unique,omitempty"`

	// TargetType is the invite's target type.
	TargetType invite.TargetType `json:"target_type,omitempty"`

	// TargetUserID is the id of the user
	// whose stream to display for the invite.
	TargetUserID snowflake.Snowflake `json:"target_user_id,omitempty"`

	// TargetAppID is the id of the embedded
	// application to open for the invite.
	TargetAppID snowflake.Snowflake `json:"target_application_id,omitempty"`
}

// CreateInvite creates an invite for
// the channel with the given id.
func CreateInvite(token auth.Token, channelID snowflake.Snowflake) (invite invite.Invite, err error) {
	return invite, client.POST(client.Request{
		Path:   fmt.Sprintf("/channels/%s/invites", channelID),
		Token:  token,
		Result: &invite,
	})
}
