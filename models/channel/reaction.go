package channel

import "github.com/gompus/gompus/models/emoji"

// A Reaction represents a reaction to a message.
// See https://discord.com/developers/docs/resources/channel#reaction-object.
type Reaction struct {
	// Count is the number of times this emoji has
	// been used to react to that particular message.
	Count int `json:"count,omitempty"`

	// Me indicates whether the current user reacted using this emoji.
	Me bool `json:"me,omitempty"`

	// Emoji is the emoji used to react to the message.
	Emoji emoji.Emoji `json:"emoji,omitempty"`
}
