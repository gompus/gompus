package appcommands

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// An OptionChoice restricts the choices for Option.
// See https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-choice-structure.
type OptionChoice struct {
	Name  string            `json:"name,omitempty"`
	Value OptionChoiceValue `json:"value,omitempty"`
}

// OptionChoiceValue represents the value of an OptionChoice.
// Valid types are int, float, string.
type OptionChoiceValue struct {
	value interface{}
}

// IsString reports whether o represents a string.
func (o OptionChoiceValue) IsString() bool {
	return o.kind() == reflect.String
}

// String returns the string representation of o.
func (o OptionChoiceValue) String() string {
	return fmt.Sprint(o.value)
}

// IsInt reports whether o represents an integer.
func (o OptionChoiceValue) IsInt() bool {
	return o.kind() == reflect.Int
}

// Int returns the integer representation of o.
// It panics if o does not represent an integer.
func (o OptionChoiceValue) Int() int {
	return o.value.(int)
}

// IsFloat reports whether o represents a floating point number.
func (o OptionChoiceValue) IsFloat() bool {
	return o.kind() == reflect.Float64
}

// Float returns the floating point representation of o.
// It panics if o does not represent a floating point number.
func (o OptionChoiceValue) Float() float64 {
	return o.value.(float64)
}

func (o OptionChoiceValue) kind() reflect.Kind {
	return reflect.TypeOf(o.value).Kind()
}

// UnmarshalJSON unmarshals data into o.
func (o *OptionChoiceValue) UnmarshalJSON(data []byte) error {
	if o.unmarshalFloat(data) || o.unmarshalInt(data) || o.unmarshalString(data) {
		return nil
	}
	return errors.New("cannot unmarshal " + string(data) + " into go value of type OptionChoiceValue: type is not string, int or float")
}

func (o *OptionChoiceValue) unmarshalString(data []byte) bool {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return false
	}
	o.value = s
	return true
}

func (o *OptionChoiceValue) unmarshalInt(data []byte) bool {
	var n int
	if err := json.Unmarshal(data, &n); err != nil {
		return false
	}
	o.value = n
	return true
}

func (o *OptionChoiceValue) unmarshalFloat(data []byte) bool {
	var n float64
	if err := json.Unmarshal(data, &n); err != nil {
		return false
	}
	o.value = n
	return true
}
