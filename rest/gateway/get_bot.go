package gateway

import (
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
)

// SessionStartLimit contains information about a session start limit.
type SessionStartLimit struct {
	// Total is the total number of session
	// starts the current user is allowed.
	Total int `json:"total,omitempty"`

	// Remaining is the remaining number of
	// session starts the current user is allowed.
	Remaining int `json:"remaining,omitempty"`

	// ResetAfter is the number of milliseconds
	// after which the session start limit resets.
	ResetAfter int `json:"reset_after,omitempty"`

	// MaxConcurrency is the number of
	// identify requests allowed per 5 seconds.
	MaxConcurrency int `json:"max_concurrency,omitempty"`
}

type GetBotResponse struct {
	// URL is the WSS URL that can be
	// used for connecting to the gateway.
	URL string `json:"url,omitempty"`

	// Shards is the recommended number
	// of shards to use when connecting.
	Shards int `json:"shards,omitempty"`

	// SessionStartLimit contains information
	// about the current session start limit.
	SessionStartLimit SessionStartLimit `json:"session_start_limit,omitempty"`
}

// GetBot retrieves information about the gateway,
// as well as additional metadata that can help
// during the operation of large or sharded bots.
func GetBot(token auth.Token) (resp GetBotResponse, err error) {
	return resp, client.GET(client.Request{
		Path:   "/gateway/bot",
		Token:  token,
		Result: &resp,
	})
}
