package guild

//go:generate stringer -type=PermiumTier

// PremiumTier represents a guil's premium tier.
type PremiumTier int

const (
	PremiumTierNone PremiumTier = iota
	PremiumTier1
	PremiumTier2
	PremiumTier3
)
