package auditlog

// Change represents a change to an audit log.
// See https://discord.com/developers/docs/resources/audit-log#audit-log-change-object.
type Change struct {
	// NewValue is the key's new value.
	NewValue interface{} `json:"new_value,omitempty"`

	// OldValue is the key's old value.
	OldValue interface{} `json:"old_value,omitempty"`

	// Key is the name of the change key.
	Key ChangeKey `json:"key,omitempty"`
}
