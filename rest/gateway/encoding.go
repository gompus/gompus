package gateway

// Encoding specified the encoding used
// for communicating over websocket.
type Encoding string

const (
	// EncodingJSON can be used to set the encoding
	// used for communicating over websocket to json.
	EncodingJSON Encoding = "json"

	// EncodingETF can be used to set the encoding
	// used for communicating over websocket to etf.
	EncodingETF Encoding = "etf"
)
