package message

import (
	"encoding/json"
	"github.com/gompus/gompus/models/application"
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/models/channel/embed"
	"github.com/gompus/gompus/models/channel/mentions"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/gompus/models/sticker"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
	"github.com/gompus/timestamp"
	"strconv"
)

// A Message represents a message sent in a Discord channel.
// See https://discord.com/developers/docs/resources/channel#message-object.
type Message struct {
	// ID is the message's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// ChannelID is the id of the channel the message was sent in.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// GuildID is the id of the guild the message was sent in.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Author is the message's author.
	Author user.User `json:"author,omitempty"`

	// Member contains member properties for the message's author.
	Member guild.Member `json:"member,omitempty"`

	// Content is the message's content.
	Content string `json:"content,omitempty"`

	// Timestamp is the time the message was sent.
	Timestamp timestamp.Timestamp `json:"timestamp,omitempty"`

	// EditedTimestamp is the time the message was last edited.
	EditedTimestamp *timestamp.Timestamp `json:"edited_timestamp,omitempty"`

	// TTS indicates whether the message is a TTS message.
	TTS bool `json:"tts,omitempty"`

	// MentionEveryone indicates whether the message mentions everyone.
	MentionEveryone bool `json:"mention_everyone,omitempty"`

	// Mentions is the set of users specifically mentioned in the message.
	Mentions []MentionedUser `json:"mentions,omitempty"`

	// MentionRoles is the set of roles specifically mentioned in the message.
	MentionRoles []snowflake.Snowflake `json:"mention_roles,omitempty"`

	// MentionChannels is the set of channels
	// specifically mentioned in the message.
	MentionChannels []mentions.Mention `json:"mention_channels,omitempty"`

	// Attachments contains any attached files.
	Attachments []channel.Attachment `json:"attachments,omitempty"`

	// Embeds contains any embedded content.
	Embeds []embed.Embed `json:"embeds,omitempty"`

	// Reactions is the set of reactions to the message.
	Reactions []channel.Reaction `json:"reactions,omitempty"`

	// Nonce is used for validating the message.
	Nonce Nonce `json:"nonce,omitempty"`

	// Pinned indicates whether the message is pinned.
	Pinned bool `json:"pinned,omitempty"`

	// WebhookID is the id of the webhook that generated the
	// message, if the message was generated by a webhook.
	WebhookID snowflake.Snowflake `json:"webhook_id,omitempty"`

	// Type is the message's type.
	Type Type `json:"type,omitempty"`

	// Activity is the message activity sent
	// with Rich Presence-related embeds.
	Activity Activity `json:"activity,omitempty"`

	// Application is the application sent
	// with Rich Presence-relatex embeds.
	Application application.Application `json:"application,omitempty"`

	// ApplicationID is the id of the interaction's application
	// if the message is a response to an interaction.
	ApplicationID snowflake.Snowflake `json:"application_id,omitempty"`

	// MessageReference shows the source of a crosspost,
	// channel follow, add, pin, or reply message.
	MessageReference Reference `json:"message_reference,omitempty"`

	// Flags is the bitwise or of the message's flags.
	Flags Flags `json:"flags,omitempty"`

	// ReferencedMessage is the message
	// associated with the message reference.
	ReferencedMessage *Message `json:"referenced_message,omitempty"`

	// Interaction is the interaction associated with the message
	// if the message is sent in response to an interaction.
	Interaction interface{} `json:"interaction,omitempty"` // todo: replace with actual type

	// Thread is the thread started from the message
	// if the message marks the beginning of a thread.
	Thread channel.Channel `json:"thread,omitempty"`

	// Components contains the message's components.
	Components []interface{} `json:"components,omitempty"` // todo: replace with actual type

	// StickerItems contains the message's stickers.
	StickerItems []sticker.Item `json:"sticker_items,omitempty"`
}

type MentionedUser struct {
	user.User
	Member guild.Member
}

type Nonce struct {
	value interface{}
}

// UnmarshalJSON unmarshals data into n.
func (n *Nonce) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	v, err := strconv.Atoi(raw)
	if err == nil {
		n.value = v
	} else {
		n.value = raw
	}
	return nil
}

// Int returns the value of n as an int.
func (n Nonce) Int() int {
	return n.value.(int)
}

// String returns the value of n as a string.
func (n Nonce) String() string {
	return n.value.(string)
}