package sticker

import "github.com/gompus/snowflake"

// Pack represents a pack of standard Stickers.
// See https://discord.com/developers/docs/resources/sticker#sticker-pack-object.
type Pack struct {
	// Id is the pack's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Stickers contains the stickers in the pack.
	Stickers []Sticker `json:"stickers,omitempty"`

	// Name is the pack's name.
	Name string `json:"name,omitempty"`

	// SkuID is the id of the pack's SKU.
	SkuID snowflake.Snowflake `json:"sku_id,omitempty"`

	// CoverStickerID is the id of a sticker in
	// the pack which is shown as the pack's icon.
	CoverStickerID snowflake.Snowflake `json:"cover_sticker_id,omitempty"`

	// Description is the pack's description.
	Description string `json:"description,omitempty"`

	// BannerAssetID is the id of the pack's banner image.
	BannerAssetID snowflake.Snowflake `json:"banner_asset_id,omitempty"`
}
