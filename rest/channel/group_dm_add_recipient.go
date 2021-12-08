package channel

import (
	"fmt"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type GroupDMAddRecipientParams struct {
	// AccessToken is the access token of the user
	// that granted the app the gdm.join scope.
	AccessToken string `json:"access_token,omitempty"`

	// Nick is the nickname of the user being added.
	Nick string `json:"nick,omitempty"`
}

// GroupDMAddRecipient adds the recipient with the
// given id to the group dm channel with the given id.
func GroupDMAddRecipient(token auth.Token, channelID, userID snowflake.Snowflake, params GroupDMAddRecipientParams) error {
	return client.PUT(client.Request{
		Path:  fmt.Sprintf("/channels/%s/recipients/%s", channelID, userID),
		Token: token,
	})
}
