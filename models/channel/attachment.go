package channel

import "github.com/gompus/snowflake"

// An Attachment is an attachment to a message.
// See https://discord.com/developers/docs/resources/channel#attachment-object.
type Attachment struct {
	// ID is the attachment's ID.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Filename is the name of the attached file.
	Filename string `json:"filename,omitempty"`

	// Description is the file's description.
	Description string `json:"description,omitempty"`

	// ContentType is the attachment's media type.
	ContentType string `json:"content_type,omitempty"`

	// Size is the file size in bytes.
	Size int `json:"size,omitempty"`

	// URL is the source url of the file.
	URL string `json:"url,omitempty"`

	// ProxyURL is a proxied url of the file.
	ProxyURL string `json:"proxy_url,omitempty"`

	// Width is the height of the file (if the file is an image).
	Width *int `json:"width,omitempty"`

	// Height is the height of the file (if the file is an image).
	Height *int `json:"height,omitempty"`

	// Ephemeral indicates whether the attachment is ephemeral.
	Ephemeral bool `json:"ephemeral,omitempty"`
}
