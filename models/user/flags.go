package user

import (
	"strings"
)

// Flags indicate a User's status.
type Flags int

const (
	// FlagsNone indicates that no flags are present.
	FlagsNone Flags = 0

	// FlagsStaff is used for Discord employees.
	FlagsStaff = 1 << 0

	// FlagsPartner is used for partnered server owners.
	FlagsPartner = 1 << 1

	// FlagsHypesquad is used for HypeSquad events coordinators.
	FlagsHypesquad = 1 << 2

	// FlagsBugHunterLevel1 is used for level 1 bug hunter.
	FlagsBugHunterLevel1 = 1 << 3

	// FlagsHypesquadOnlineHouse1 is used for house bravery members.
	FlagsHypesquadOnlineHouse1 = 1 << 6

	// FlagsHypesquadOnlineHouse2 is used for house brilliance members.
	FlagsHypesquadOnlineHouse2 = 1 << 7

	// FlagsHypesquadOnlineHouse3 is used for house balance members.
	FlagsHypesquadOnlineHouse3 = 1 << 8

	// FlagsPremiumEarlySupporter is used for early Nitro supporters.
	FlagsPremiumEarlySupporter = 1 << 9

	// FlagsTeamPseudoUser is used for users representing a team.
	FlagsTeamPseudoUser = 1 << 10

	// FlagsBugHunterLevel2 is used for level 2 bug hunters.
	FlagsBugHunterLevel2 = 1 << 14

	// FlagsVerifiedBot is used for verified bots.
	FlagsVerifiedBot = 1 << 16

	// FlagsVerifiedDeveloper is used for early verified bot developers.
	FlagsVerifiedDeveloper = 1 << 17

	// FlagsCertifiedModerator is used for certified Discord moderators.
	FlagsCertifiedModerator = 1 << 18

	// FlagsBotHttpInteractions is used for bots that only use HTTP interactions.
	FlagsBotHttpInteractions = 1 << 19
)

var flagToString = map[Flags]string{
	FlagsNone:                  "None",
	FlagsStaff:                 "Staff",
	FlagsPartner:               "Partner",
	FlagsHypesquad:             "Hypesquad",
	FlagsBugHunterLevel1:       "BugHunterLevel1",
	FlagsHypesquadOnlineHouse1: "HypesquadOnlineHouse1",
	FlagsHypesquadOnlineHouse2: "HypesquadOnlineHouse2",
	FlagsHypesquadOnlineHouse3: "HypesquadOnlineHouse3",
	FlagsPremiumEarlySupporter: "PremiumEarlySupporter",
	FlagsTeamPseudoUser:        "TeamPseudoUser",
	FlagsBugHunterLevel2:       "BugHunterLevel2",
	FlagsVerifiedBot:           "VerifiedBot",
	FlagsVerifiedDeveloper:     "VerifiedDeveloper",
	FlagsCertifiedModerator:    "CertifiedModerator",
	FlagsBotHttpInteractions:   "BotHttpInteractions",
}

// String returns the string representation of f.
func (f Flags) String() string {
	var out []string
	for flag, s := range flagToString {
		if f&flag != 0 {
			out = append(out, s)
		}
	}
	return strings.Join(out, "|")
}
