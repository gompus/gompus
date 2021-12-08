package sticker

import (
	"github.com/gompus/gompus/models/sticker"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
)

// GetNitroPacks retrieves a set of sticker
// packs available to Nitro subscribers.
func GetNitroPacks(token auth.Token) (packs []sticker.Pack, err error) {
	resp := struct {
		Packs []sticker.Pack `json:"sticker_packs"`
	}{}
	err = client.GET(client.Request{
		Path:   "/sticker-packs",
		Token:  token,
		Result: &resp,
	})
	return resp.Packs, err
}
