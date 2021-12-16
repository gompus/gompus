package commands

// IdentifyProperties contain client information.
// See https://discord.com/developers/docs/topics/gateway#identify-identify-connection-properties.
type IdentifyProperties struct {
	// OS is the client's os.
	OS string `json:"$os,omitempty"`

	// Browser is the client's browser.
	Browser string `json:"$browser,omitempty"`

	// Device is an identifier of the client's device.
	Device string `json:"$device,omitempty"`
}
