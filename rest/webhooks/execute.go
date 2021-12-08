package webhooks

import (
	"fmt"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/models/channel/embed"
	"github.com/gompus/gompus/models/channel/mentions"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/snowflake"
)

type ExecuteParams struct {
	// Wait indicates whether to wait for server
	// confirmation of the message before responding.
	Wait bool

	// ThreadID can be used to specifiy the id of a thread
	// within a webhook's channel a message should be sent to.
	ThreadID snowflake.Snowflake

	// Content is the message's content.
	Content string

	// Username can be used to override a
	// webhook's default username.
	Username string

	// AvatarURL can be used to override
	// a webhook's default avatar.
	AvatarURL string

	// TTS indicates whether a TTS message is being sent.
	TTS bool

	// Embeds contains embedded rich content.
	Embeds []embed.Embed

	// AllowedMentions specifies allowed
	// mentions for a message.
	AllowedMentions mentions.Allowed

	// Components contains components
	// to include with the message.
	Components []interface{}

	// Files contains files to include with the message.
	Files []byte

	// Payload contains the JSON encoded body of non-file params.
	Payload string

	// Attachments contains attachments
	// to include with the message.
	Attachments []channel.Attachment
}

func (p ExecuteParams) queryParams() executeQueryParams {
	return executeQueryParams{
		Wait:     p.Wait,
		ThreadID: p.ThreadID,
	}
}

type executeQueryParams struct {
	Wait     bool                `json:"wait,omitempty"`
	ThreadID snowflake.Snowflake `json:"thread_id,omitempty"`
}

func (p ExecuteParams) jsonParams() executeJsonParams {
	return executeJsonParams{
		Content:         p.Content,
		Username:        p.Username,
		AvatarURL:       p.AvatarURL,
		TTS:             p.TTS,
		Embeds:          p.Embeds,
		AllowedMentions: p.AllowedMentions,
		Components:      p.Components,
		Files:           p.Files,
		Payload:         p.Payload,
		Attachments:     p.Attachments,
	}
}

type executeJsonParams struct {
	Content         string               `json:"content,omitempty"`
	Username        string               `json:"username,omitempty"`
	AvatarURL       string               `json:"avatar_url,omitempty"`
	TTS             bool                 `json:"tts,omitempty"`
	Embeds          []embed.Embed        `json:"embeds,omitempty"`
	AllowedMentions mentions.Allowed     `json:"allowed_mentions,omitempty"`
	Components      interface{}          `json:"components,omitempty"`
	Files           []byte               `json:"files,omitempty"`
	Payload         string               `json:"payload_json,omitempty"`
	Attachments     []channel.Attachment `json:"attachments,omitempty"`
}

// Execute executes the webhook with the
// given id, using the given webhook token.
func Execute(token string, hookID snowflake.Snowflake, params ExecuteParams) error {
	return client.POST(client.Request{
		Path:  fmt.Sprintf("/webhooks/%s/%s", hookID, token),
		Query: client.GenerateQuery(params.queryParams()),
		Body:  params.jsonParams(),
	})
}

type ExecuteSlackParams struct {
	// ThreadID is the id of the thread to send the message in.
	ThreadID snowflake.Snowflake `json:"thread_id,omitempty"`

	// Wait indicates whether to wait for server
	// confirmation of the message before responding.
	Wait bool `json:"wait,omitempty"`
}

// ExecuteSlack executes the slack-compatible webhook
// with the given id, using the given webhook token.
func ExecuteSlack(token string, hookID snowflake.Snowflake, params ...ExecuteParams) error {
	return client.POST(client.Request{
		Path:  fmt.Sprintf("/webhooks/%s/%s/slack", hookID, token),
		Query: client.GenerateQuery(params),
	})
}

type ExecuteGithubParams struct {
	// ThreadID is the id of the thread to send the message in.
	ThreadID snowflake.Snowflake `json:"thread_id,omitempty"`

	// Wait indicates whether to wait for server
	// confirmation of the message before responding.
	Wait bool `json:"wait,omitempty"`
}

// ExecuteGithub executes the github-compatible webhook
// with the given id, using the given webhook token.
func ExecuteGithub(token string, hookID snowflake.Snowflake, params ...ExecuteGithubParams) error {
	return client.POST(client.Request{
		Path:  fmt.Sprintf("/webhooks/%s/%s/slack", hookID, token),
		Query: client.GenerateQuery(params),
	})
}
