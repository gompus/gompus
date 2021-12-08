package client

import "reflect"

// MergeParams merges the given params. It expects a slice
// of values and returns the first of those values. If the
// slice is empty, MergeParams returns nil.
func MergeParams(params interface{}) interface{} {
	if params == nil {
		panic("cannot merge nil params")
	}
	
	rt := reflect.TypeOf(params)
	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		rv := reflect.ValueOf(params)
		if rv.Len() > 0 {
			return rv.Index(0)
		}
		return nil
	default:
		return params
	}
}
