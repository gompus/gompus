package embed

// Video represents an embedded video.
// See https://discord.com/developers/docs/resources/channel#embed-object-embed-video-structure.
type Video struct {
	// URL is the source url of the video.
	URL string `json:"url,omitempty"`

	// ProxyURL is a proxied url of the video.
	ProxyURL string `json:"proxy_url,omitempty"`

	// Width is the width of the video.
	Width int `json:"width,omitempty"`

	// Height is the height of the video.
	Height int `json:"height,omitempty"`
}
