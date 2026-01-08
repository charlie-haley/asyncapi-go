package http

// OperationBinding represents the HTTP Operation Binding object.
//
// This object contains information about the operation representation in HTTP.
// +binding
type OperationBinding struct {
	// Method is the HTTP method for the operation.
	// MUST be one of GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS, CONNECT, or TRACE.
	Method string `json:"method,omitempty"`

	// Query contains the definitions of the query parameters.
	// This schema MUST be of type object and have a properties key.
	Query interface{} `json:"query,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
