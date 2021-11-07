package application

import (
	"github.com/gompus/gompus/models/team"
	"github.com/gompus/gompus/models/user"
	"github.com/gompus/snowflake"
)

// An Application represents a Discord application.
// See https://discord.com/developers/docs/resources/application#application-resource.
type Application struct {
	// ID is the app's id.
	ID snowflake.Snowflake `json:"id,omitempty"`

	// Name is the app's name.
	Name string `json:"name,omitempty"`

	// Icon is the app's icon hash.
	Icon *string `json:"icon,omitempty"`

	// Description is the app's description.
	Description string `json:"description,omitempty"`

	// RpcOrigins is a list of rpc origin urls, if rpc is enabled.
	RpcOrigins []string `json:"rpc_origins,omitempty"`

	// BotPublic indicates whether everyone can join the app's
	// bot to guilds (if false, only the app's owner can do so).
	BotPublic bool `json:"bot_public,omitempty"`

	// BotRequireCodeGrant indicates whether the app's bot will only
	// join upon completion of the full oauth2 code grant flow.
	BotRequireCodeGrant bool `json:"bot_require_code_grant,omitempty"`

	// TermsOfServiceURL is the URL of the app's terms of service.
	TermsOfServiceURL string `json:"terms_of_service_url,omitempty"`

	// PrivacyPolicyURL is the URL of the app's privacy policy.
	PrivacyPolicyURL string `json:"privacy_policy_url,omitempty"`

	// Owner is a partial user containing info on the app's owner.
	Owner user.User `json:"owner,omitempty"`

	// Summary is the summary field for the store page of the
	// app's primary sku if the app is a game sold on Discord.
	Summary string `json:"summary,omitempty"`

	// VerifyKey is the hex encoded key for verification
	// in interactions and the GameSDK's get GetTicket.
	VerifyKey string `json:"verify_key,omitempty"`

	// Team contains a list of the members of the
	// app's team if the app belongs to a team.
	Team *team.Team `json:"team,omitempty"`

	// GuildID is the id of the guild to which the app has
	// been linked if the app is a game sold on Discord.
	GuildID snowflake.Snowflake `json:"guild_id,omitempty"`

	// PrimarySkuID is the id of the "Game SKU" that
	// is created if the app is a game sold on Discord.
	PrimarySkuID snowflake.Snowflake `json:"primary_sku_id,omitempty"`

	// Slug is the URL slug that links to the app's
	// store page if the app is a game sold on Discord.
	Slug string `json:"slug,omitempty"`

	// CoverImage is the app's default rich
	// presence invite cover image hash.
	CoverImage string `json:"cover_image,omitempty"`

	// Flags contains the app's public Flags.
	Flags Flags `json:"flags,omitempty"`
}
