package client

import "net/http"

// Method is an explicit type for an HTTP method.
type Method string

const (
	Get     Method = http.MethodGet
	Post    Method = http.MethodPost
	Put     Method = http.MethodPut
	Patch   Method = http.MethodPatch
	Delete  Method = http.MethodDelete
)
