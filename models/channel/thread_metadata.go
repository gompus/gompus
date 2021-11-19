package channel

import "github.com/gompus/timestamp"

// ThreadMetadata contains thread-specific channel fields.
// See https://discord.com/developers/docs/resources/channel#thread-metadata-object.
type ThreadMetadata struct {
	// Archived indicates whether the thread is archived.
	Archived bool `json:"archived,omitempty"`

	// AutoArchiveDuration is the duration after
	// which the thread is to be archived.
	AutoArchiveDuration int `json:"auto_archive_duration,omitempty"`

	// ArchiveTimestamp represents the point in time when
	// the thread's archive status was last changed.
	ArchiveTimestamp timestamp.Timestamp `json:"archive_timestamp,omitempty"`

	// Locked indicates whether the thread is locked.
	Locked bool `json:"locked,omitempty"`

	// Invitable indicates whether non-moderators
	// can add other non-moderators to the thread.
	Invitable bool `json:"invitable,omitempty"`
}
