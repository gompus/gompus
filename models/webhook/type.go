package webhook

//go:generate stringer -type=Type -trimprefix=Type

// Type represents the type of a Webhook.
type Type int

const (
	// TypeIncoming is the type for webhooks that can
	// post messages to channels with a generated token.
	TypeIncoming Type = iota + 1

	// TypeChannelFollower is the type for internal webhooks
	// used with channel following to post new messages to channels.
	TypeChannelFollower

	// TypeApplication is the type for webhooks used with interactions.
	TypeApplication
)
