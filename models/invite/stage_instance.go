package invite

import "github.com/gompus/gompus/models/guild"

// StageInstance represents an invite to a stage.
// See https://discord.com/developers/docs/resources/invite#invite-stage-instance-object.
type StageInstance struct {
	// Members contains the members speaking on the stage.
	Members []guild.Member `json:"members,omitempty"`

	// ParticipationCount is the number of users on the stage.
	ParticipantCount int `json:"participant_count,omitempty"`

	// SpeakerCount is the number of users speaking on the stage.
	Speakercount int `json:"speakercount,omitempty"`

	// Topic is the stage's topic.
	Topic string `json:"topic,omitempty"`
}
