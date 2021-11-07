package user

// A Connection contains information about a User's connection.
// See https://discord.com/developers/docs/resources/user#connection-object.
type Connection struct {
	// ID is the if of the connected account.
	ID string `json:"id,omitempty"`

	// Name is the connected account's username.
	Name string `json:"name,omitempty"`

	// Type represents the connection's type (e.g. "twitch", "youtube").
	Type string `json:"type,omitempty"`

	// Revoked indicates whether the connection is revoked.
	Revoked bool `json:"revoked,omitempty"`

	// Integrations contains partial server integrations.
	Integrations []interface{} `json:"integrations,omitempty"` // todo: replace with actual type

	// Verified indicates whether the connection is verified.
	Verified bool `json:"verified,omitempty"`

	// FriendSync indicates whether friend sync is enabled for this connection.
	FriendSync bool `json:"friend_sync,omitempty"`

	// ShowActivity indicates whether activities relates to
	// this connection will be shown in presence updates.
	ShowActivity bool `json:"show_activity,omitempty"`

	// Visibility represents the visibility of this connection.
	Visibility Visibility `json:"visibility,omitempty"`
}
