package team

//go:generate stringer -type=MembershipState -trimprefix=MembershipState

// MembershipState represents the state of a team membership.
type MembershipState int

const (
	MembershipStateInvited  MembershipState = 1
	MembershipStateAccepted                 = 2
)
