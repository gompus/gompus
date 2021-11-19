package embed

// Image represents an embedded image.
// See https://discord.com/developers/docs/resources/channel#embed-object-embed-image-structure.
type Image struct {
	// URL is the source url of the image.
	URL string `json:"url,omitempty"`

	// ProxyURL is a proxied url of the image.
	ProxyURL string `json:"proxy_url,omitempty"`

	// Width is the width of the image.
	Width int `json:"width,omitempty"`

	// Height is the height of the image.
	Height int `json:"height,omitempty"`
}
