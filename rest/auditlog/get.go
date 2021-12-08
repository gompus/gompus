package auditlog

import (
	"fmt"
	"github.com/gompus/gompus/models/auditlog"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

type GetParams struct {
	// UserID can be used to filter the log
	// for actions made by a specific user.
	UserID snowflake.Snowflake `json:"user_id,omitempty"`

	// ActionType can be used to filter
	// the log for specific event types.
	ActionType int `json:"action_type,omitempty"`

	// Before can be used to filter
	// the log for a certain entry id.
	Before snowflake.Snowflake `json:"before,omitempty"`

	// Limit is the maximum number of entries to retrieve.
	Limit int `json:"limit,omitempty"`
}

// Get retrieves the audit log of the guild with the given id.
func Get(token auth.Token, guildID snowflake.Snowflake, params ...GetParams) (auditlog auditlog.AuditLog, err error) {
	return auditlog, client.GET(client.Request{
		Path:   fmt.Sprintf("/guilds/%s/audit-logs", guildID),
		Query:  client.GenerateQuery(params),
		Token:  token,
		Result: &auditlog,
	})
}
