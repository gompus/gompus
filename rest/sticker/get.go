package sticker

import (
	"fmt"
	"github.com/gompus/gompus/models/sticker"
	"github.com/gompus/gompus/rest/client"
	"github.com/gompus/gompus/rest/client/auth"
	"github.com/gompus/snowflake"
)

// Get retrieves the sticker with the given id.
func Get(token auth.Token, id snowflake.Snowflake) (sticker sticker.Sticker, err error) {
	return sticker, client.GET(client.Request{
		Path:   fmt.Sprintf("/stickers/%s", id),
		Token:  token,
		Result: &sticker,
	})
}
