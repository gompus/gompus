package embed

// Thumbnail represents an embedded thumbnail.
// See https://discord.com/developers/docs/resources/channel#embed-object-embed-thumbnail-structure.
type Thumbnail struct {
	// URL is the source URL of the thumbnail.
	URL string `json:"url,omitempty"`

	// ProxyURL a the proxied URL of the thumbnail.
	ProxyURL string `json:"proxy_url,omitempty"`

	// Width is the width of the thumbnail.
	Width int `json:"width,omitempty"`

	// Height is the height of the thumbnail.
	Height int `json:"height,omitempty"`
}
