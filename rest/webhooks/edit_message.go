package webhooks

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/models/channel/embed"
	"github.com/gompus/gompus/models/channel/mentions"
	"github.com/gompus/gompus/models/channel/message"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/snowflake"
)

type EditMessageParams struct {
	// ThreadID is the id of the thread the message is in.
	ThreadID snowflake.Snowflake

	// Content is the message content.
	Content string

	// Embeds contains embedded rich content.
	Embeds []embed.Embed

	// AllowedMentions are the message's allowed mentions.
	AllowedMentions mentions.Allowed

	// Components contains the components to include with the message.
	Components []interface{}

	// Files contains the files to include with the message.
	Files []byte

	// Payload contains the JSON encoded body of non-file params.
	Payload string

	// Attachments contains attached files.
	Attachments []channel.Attachment
}

func (p EditMessageParams) queryParams() editMessageQueryParams {
	return editMessageQueryParams{
		ThreadID: p.ThreadID,
	}
}

type editMessageQueryParams struct {
	ThreadID snowflake.Snowflake `json:"thread_id,omitempty"`
}

func (p EditMessageParams) jsonParams() editMessageJsonParams {
	return editMessageJsonParams{
		Content:         p.Content,
		Embeds:          p.Embeds,
		AllowedMentions: p.AllowedMentions,
		Components:      p.Components,
		Files:           p.Files,
		Payload:         p.Payload,
		Attachments:     p.Attachments,
	}
}

type editMessageJsonParams struct {
	Content         string               `json:"content,omitempty"`
	Embeds          []embed.Embed        `json:"embeds,omitempty"`
	AllowedMentions mentions.Allowed     `json:"allowed_mentions,omitempty"`
	Components      []interface{}        `json:"components,omitempty"`
	Files           []byte               `json:"files,omitempty"`
	Payload         string               `json:"payload_json,omitempty"`
	Attachments     []channel.Attachment `json:"attachments,omitempty"`
}

// EditMessage edits the message with the given id that was
// sent on the webhook with the given id.
func EditMessage(token string, hookID, messageID snowflake.Snowflake, params EditMessageParams) (msg message.Message, err error) {
	return msg, client.PATCH(client.Request{
		Path:  fmt.Sprintf("/webhhoks/%s/%s/messages/%s", hookID, token, messageID),
		Query: client.GenerateQuery(params.queryParams()),
		Body:  params.jsonParams(),
	})
}
