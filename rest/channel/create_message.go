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

type CreateMessageParams struct {
	// Content is the message's content.
	Content string `json:"content,omitempty"`

	// TTS indicates whether the message is a TTS message.
	TTS bool `json:"tts,omitempty"`

	// Embeds contains the message's embeds.
	Embeds []embed.Embed `json:"embeds,omitempty"`

	// AllowedMentions are the message's allowed mentions.
	AllowedMentions mentions.Allowed `json:"allowed_mentions,omitempty"`

	// MessageReference is a reference to another message.
	MessageReference message.Reference `json:"message_reference,omitempty"`

	// Components contains the message's components.
	Components []interface{} `json:"components,omitempty"`

	// StickerIDs contains the ids of up
	// to 3 stickers to send in the message.
	StickerIDs []snowflake.Snowflake `json:"sticker_i_ds,omitempty"`

	// Files contains the files to send with the message.
	Files []byte `json:"files,omitempty"`

	// Payload is the JSON encoded body of non-file params.
	Payload string `json:"payload_json,omitempty"`

	// Attachments contains the message's attachments.
	Attachments []channel.Attachment `json:"attachments,omitempty"`
}

// CreateMessage creates a new message in the channel with the given id.
func CreateMessage(token auth.Token, channelID snowflake.Snowflake, params CreateMessageParams) (msg message.Message, err error) {
	return msg, client.POST(client.Request{
		Path:   fmt.Sprintf("/channels/%s/messages", channelID),
		Body:   params,
		Token:  token,
		Result: &msg,
	})
}
