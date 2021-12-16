package events

import "github.com/gompus/gompus/models/stage"

// StageInstanceCreateEvent is the event
// sent when a stage instance is created.
type StageInstanceCreateEvent stage.Stage

// StageInstanceUpdateEvent is the event
// sent when a stage instance is updated.
type StageInstanceUpdateEvent stage.Stage

// StageInstanceDeleteEvent is the event
// sent when a stage instance is deleted.
type StageInstanceDeleteEvent stage.Stage
