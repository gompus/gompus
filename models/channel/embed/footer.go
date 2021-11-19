package embed

// Footer represents an embed's footer.
// See https://discord.com/developers/docs/resources/channel#embed-object-embed-footer-structure.
type Footer struct {
	// Text is the footer text.
	Text string `json:"text,omitempty"`

	// IconURL is the url of the footer icon.
	IconURL string `json:"icon_url,omitempty"`

	// ProxyIconURL is a proxied url of the footer icon.
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}
