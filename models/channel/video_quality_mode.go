package channel

//go:generate stringer -type=VideoQualityMode -trimprefix=VideoQualityMode

// VideoQualityMode describes the video quality of a channel.
type VideoQualityMode int

const (
	VideoQualityModeAuto VideoQualityMode = 1
	VideoQualityModeFull
)
