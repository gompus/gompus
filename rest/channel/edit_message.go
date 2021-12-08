package channel

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/models/channel/embed"
	"github.com/gompus/gompus/models/channel/mentions"
	"github.com/gompus/gompus/models/channel/message"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type EditMessageParams struct {
	// Content is the message's content.
	Content string `json:"content,omitempty"`

	// Embeds contains the message's embedded rich content.
	Embeds []embed.Embed `json:"embeds,omitempty"`

	// Flags contains the message's flags.
	Flags message.Flags `json:"flags,omitempty"`

	// Mentions contains the message's allowed mentions.
	AllowedMentions mentions.Allowed `json:"allowed_mentions,omitempty"`

	// Componenets contains the message's components.
	Components []interface{} `json:"components,omitempty"`

	// Files contains the contents of the message's attached file.
	Files []byte `json:"files,omitempty"`

	// Payload contains the JSON encoded body of non-file params.
	Payload string `json:"payload_json,omitempty"`

	// Attachment contains the message's attachment.
	Attachments []channel.Attachment `json:"attachments,omitempty"`
}

// EditMessage edits the message with the given
// id in the channel with the given id.
func EditMessage(token auth.Token, channelID, messageID snowflake.Snowflake, params ...EditMessageParams) (msg message.Message, err error) {
	return msg, client.PATCH(client.Request{
		Path:   fmt.Sprintf("/channels/%s/messages/%s", channelID, messageID),
		Body:   client.MergeParams(params),
		Token:  token,
		Result: &msg,
	})
}
