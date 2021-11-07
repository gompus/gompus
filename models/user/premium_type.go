package user

//go:generate stringer -type=PremiumType -trimprefix=PremiumType

// PremiumType is the type of a User's Nitro subscription.
type PremiumType int

const (
	// PremiumTypeNone indicates no Nitro subscription.
	PremiumTypeNone PremiumType = iota

	// PremiumTypeNitroClassic indicates a Nitro Classic subscription.
	PremiumTypeNitroClassic

	// PremiumTypeNitro indicates a Nitro subscription.
	PremiumTypeNitro
)
