package payloads

import (
	"encoding/json"
	"reflect"
)

// Payload represents a Discord gateway payload.
// See https://discord.com/developers/docs/topics/gateway#payloads.
type Payload struct {
	// Op is the payload's opcode.
	Op Op `json:"op,omitempty"`

	// Data is the payload's data.
	Data *json.RawMessage `json:"d,omitempty"`

	// SeqNum is the payload's sequence number.
	SeqNum *int `json:"s,omitempty"`

	// Name is the payload's name.
	Name *string `json:"t,omitempty"`
}

// DecodeData decodes p's data into v.
func (p Payload) DecodeData(v interface{}) error {
	return json.Unmarshal(*p.Data, v)
}

// DecodeDataIntoField decodes p's data into the first field of v.
func (p Payload) DecodeDataIntoField(v interface{}) error {
	field := reflect.New(reflect.TypeOf(v).Field(0).Type).Interface()
	if err := json.Unmarshal(*p.Data, field); err != nil {
		return err
	}

	reflect.ValueOf(v).Field(0).Set(reflect.ValueOf(field))
	return nil
}
