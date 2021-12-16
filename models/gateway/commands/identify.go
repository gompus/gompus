package commands

import "github.com/gompus/gompus/models/gateway"

// Identify is the command used to trigger
// initial handshakes with the gateway.
// See https://discord.com/developers/docs/topics/gateway#identify.
type Identify struct {
	// Token is the authentication token.
	Token string `json:"token,omitempty"`

	// Properties are the connection properties.
	Properties IdentifyProperties `json:"properties,omitempty"`

	// Compress indicates whether the connection supports compression.
	Compress bool `json:"compress,omitempty"`

	// LargeThreshold is the largest number of members
	// such that offline members will be included in
	// guild members lists.
	LargeThreshold int `json:"large_threshold,omitempty"`

	// Shard contains information for guild sharding.
	Shard [2]int `json:"shard,omitempty"`

	// Presence contains initial presence information.
	Presence UpdatePresence `json:"presence,omitempty"`

	// Intents contains the intents for which the
	// client wants to receive gateway events.
	Intents gateway.Intents `json:"intents,omitempty"`
}
