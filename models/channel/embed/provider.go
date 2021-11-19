package embed

// Provider contains information about the provider of an embed.
// See https://discord.com/developers/docs/resources/channel#embed-object-embed-provider-structure.
type Provider struct {
	// Name is the provider's name.
	Name string `json:"name,omitempty"`

	// URL is the provider's url.
	URL string `json:"url,omitempty"`
}
