package integration

// Account represents an Integration account.
// See https://discord.com/developers/docs/resources/guild#integration-account-object.
type Account struct {
	// ID is the account's id.
	ID string `json:"id,omitempty"`

	// Name is the account's name.
	Name string `json:"name,omitempty"`
}
