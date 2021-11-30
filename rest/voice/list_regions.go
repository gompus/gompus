package voice

import (
	"github.com/gompus/gompus/models/voice"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
)

// ListRegions retrieves a set of voice regions that can be
// used when setting a voice or stage channel's rtc region.
func ListRegions(token auth.Token) (regions []voice.Region, err error) {
	return regions, client.GET(client.Request{
		Path: "/voice/regions",
		Headers: []client.HeaderFunc{
			client.AuthHeader(token),
		},
		Result: &regions,
	})
}
