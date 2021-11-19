package embed

import "github.com/gompus/timestamp"

// An Embed represents a message embed.
// See https://discord.com/developers/docs/resources/channel#embed-object.
type Embed struct {
	// Title is the embed's title.
	Title string `json:"title,omitempty"`

	// Type is the embed's type.
	Type Type `json:"type,omitempty"`

	// Description is the embed's description.
	Description string `json:"description,omitempty"`

	// URL is the embed's URL.
	URL string `json:"url,omitempty"`

	// Timestamp is the timestamp of the embed's content.
	Timestamp timestamp.Timestamp `json:"timestamp,omitempty"`

	// Color is the embed's color.
	Color int `json:"color,omitempty"`

	// Footer contains information about the embed's footer.
	Footer Footer `json:"footer,omitempty"`

	// Image contains information about the embed's image.
	Image Image `json:"image,omitempty"`

	// Thumbnail contains information about the embed's thumbnail.
	Thumbnail Thumbnail `json:"thumbnail,omitempty"`

	// Video contains information about the embed's video.
	Video Video `json:"video,omitempty"`

	// Provider contains information about the embed's provider.
	Provider Provider `json:"provider,omitempty"`

	// Author contains information about the embed's author.
	Author Author `json:"author,omitempty"`

	// Fields contains information about the embed's fields.
	Fields []Field `json:"fields,omitempty"`
}
