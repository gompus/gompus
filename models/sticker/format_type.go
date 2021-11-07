package sticker

//go:generate stringer -type=FormatType -trimprefix=FormatType

// FormatType represents a Sticker's format.
type FormatType int

const (
	FormatTypePNG FormatType = iota + 1
	FormatTypeAPNG
	FormatTypeLottie
)
