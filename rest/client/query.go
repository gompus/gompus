package client

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

const tagKey = "json"

// GenerateForm generates a URL encoded form from v.
func GenerateForm(v interface{}) string {
	return GenerateQuery(v).Encode()
}

// GenerateQuery generates a URL query from v.
func GenerateQuery(v interface{}) url.Values {
	values := make(url.Values)
	iterStructFields(v, stringifyField(values))
	return values
}

func iterStructFields(v interface{}, fn func(field reflect.StructField, val reflect.Value)) {
	rv := reflect.Indirect(reflect.ValueOf(v))
	rt := rv.Type()

	if rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array {
		if rv.Len() == 0 {
			return
		}

		rv = rv.Index(0)
		rt = rv.Type()

		if rv.IsZero() {
			return
		}
	}

	if rt.Kind() == reflect.Struct {
		for i := 0; i < rt.NumField(); i++ {
			fn(rt.Field(i), rv.Field(i))
		}
	}
}

type values interface {
	Set(key, value string)
}

func stringifyField(target values) func(field reflect.StructField, value reflect.Value) {
	return func(field reflect.StructField, value reflect.Value) {
		if !shouldOmit(value) {
			target.Set(getFieldName(field), fmt.Sprint(reflect.Indirect(value)))
		}
	}
}

func shouldOmit(value reflect.Value) bool {
	if value.Kind() == reflect.Ptr && value.IsNil() {
		return true
	}
	return value.IsZero()
}

func getFieldName(field reflect.StructField) string {
	name, ok := field.Tag.Lookup(tagKey)
	if ok {
		return strings.Split(name, ",")[0]
	}
	return field.Name
}
