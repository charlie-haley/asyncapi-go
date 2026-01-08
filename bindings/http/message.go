package http

// MessageBinding represents the HTTP Message Binding object.
//
// This object contains information about the message representation in HTTP.
// +binding
type MessageBinding struct {
	// Headers contains the definitions of the HTTP headers.
	// This schema MUST be of type object and have a properties key.
	Headers interface{} `json:"headers,omitempty"`

	// StatusCode is the HTTP response status code according to RFC 9110.
	// Only relevant for messages referenced by the Operation Reply Object.
	StatusCode int `json:"statusCode,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
