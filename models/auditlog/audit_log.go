package auditlog

import (
	"github.com/gompus/gompus/models/channel"
	"github.com/gompus/gompus/models/guild/integration"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/gompus/models/webhook"
)

// AuditLog represents a log of admin actions performed on a guild.
// See https://discord.com/developers/docs/resources/audit-log#audit-logs-resource.
type AuditLog struct {
	// Entries contains the audit log's entries.
	Entries []Entry `json:"audit_log_entries,omitempty"`

	// Integrations contains the audit log's threads.
	Integrations []integration.Integration `json:"integrations,omitempty"`

	// Threads contains the audit log's threads.
	Threads []channel.Channel `json:"threads,omitempty"`

	// Users contains the audit log's users.
	Users []user.User `json:"users,omitempty"`

	// Webhooks contains the audit log's webhooks.
	Webhooks []webhook.Webhook `json:"webhooks,omitempty"`
}
