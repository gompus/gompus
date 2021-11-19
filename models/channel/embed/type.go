package embed

// Type is the type of an Embed.
type Type string

const (
	TypeRich    Type = "rich"
	TypeImage   Type = "image"
	TypeVideo   Type = "video"
	TypeGifv    Type = "gifv"
	TypeArticle Type = "article"
	TypeLink    Type = "link"
)
