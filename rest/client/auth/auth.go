package auth

// Token represents a Discord authentication token.
type Token string

// NewBotToken creates a new bot Token from the given token.
func NewBotToken(token string) Token {
	return Token("Bot " + token)
}

// NewBearerToken creates a new bearer Token from the given token.
func NewBearerToken(token string) Token {
	return Token("Bearer " + token)
}
