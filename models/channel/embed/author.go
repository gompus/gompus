package embed

// Author represents the author of an embed.
// See https://discord.com/developers/docs/resources/channel#embed-object-embed-author-structure.
type Author struct {
	// Name is the author's name.
	Name string `json:"name,omitempty"`

	// URL is the author's URL.
	URL string `json:"url,omitempty"`

	// Icon is the url of the author's icon.
	IconURL string `json:"icon_url,omitempty"`

	// ProxyIconURL is a proxied url of the author's icon.
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}
