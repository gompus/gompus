package events

import (
	"github.com/gompus/gompus/models/channel/message"
	"github.com/gompus/gompus/models/emoji"
	"github.com/gompus/gompus/models/guild"
	"github.com/gompus/snowflake"
)

// MessageCreateEvent is the event sent when a message is created.
type MessageCreateEvent message.Message

// MessageUpdateEvent is the event sent when a message is updated.
type MessageUpdateEvent message.Message

// MessageDeleteEvent is the event sent when a message is deleted.
type MessageDeleteEvent struct {
	// ID is the id of the deleted message.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// ChannelID is the id of the channel the message was in.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// GuildID is the id of the guild the channel is in.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`
}

// MessageDeleteBulkEvent is the event sent
// when multiple messages are deleted at once.
type MessageDeleteBulkEvent struct {
	// IDs contains the messages' ids.
	IDs []snowflake.Snowflake `json:"ids,omitempty"`

	// ChannelID is the id of the channel the messages were in.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// GuildID is the id of the guild the channel is in.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`
}

// MessageReactionAddEvent is the event sent
// when a user adds a reaction to a message.
type MessageReactionAddEvent struct {
	// UserID is the id of the user who added the reaction.
	UserID snowflake.Snowflake `json:"user_id,omitempty"`

	// ChannelID is the id of the channel the message was in.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// MessageID is the message's id.
	MessageID snowflake.Snowflake `json:"message_id,omitempty"`

	// GuildID is the id of the guild the channel is in.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Member is the member who reacted to the message
	// (present if the reaction happened in a guild).
	Member guild.Member `json:"member,omitempty"`

	// Emoji is the emoji with which the user reacted.
	Emoji emoji.Emoji `json:"emoji,omitempty"`
}

// MessageReactionRemoveEvent is the event sent
// when a user removes a reaction from a message.
type MessageReactionRemoveEvent struct {
	// UserID is the id of the user who removed the reaction.
	UserID snowflake.Snowflake `json:"user_id,omitempty"`

	// ChannelID is the id of the channel the message is in.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// MessageID is the message's id.
	MessageID snowflake.Snowflake `json:"message_id,omitempty"`

	// GuildID is the id of the guild the channel is in.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// Emoji is the emoji for which the user removed the reaction.
	Emoji emoji.Emoji `json:"emoji,omitempty"`
}

// MessageReactionRemoveAllEvent is the event sent when
// a user explicitly removes all reactions from a message.
type MessageReactionRemoveAllEvent struct {
	// ChannelID is the id of the channel the message is in.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// MessageID is the message's id.
	MessageID snowflake.Snowflake `json:"message_id,omitempty"`

	// GuildID is the id of the guild the channel is in.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`
}

// MessageReactionRemoveEmojiEvent is the event sent when a bot
// removes all instances of a given emoji from the reactions of
// a message.
type MessageReactionRemoveEmojiEvent struct {
	// ChannelID is the id of the channel the message is in.
	ChannelID snowflake.Snowflake `json:"channel_id,omitempty"`

	// GuildID is the id of the guild the channel is in.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// MessageID is the message's id.
	MessageID snowflake.Snowflake `json:"message_id,omitempty"`

	// Emoji is the removed emoji.
	Emoji emoji.Emoji `json:"emoji,omitempty"`
}
