package websockets

// ChannelBinding represents the WebSockets Channel Binding object.
//
// When using WebSockets, the channel represents the connection.
// +binding
type ChannelBinding struct {
	// Method is the HTTP method to use when establishing the connection.
	// MUST be either GET or POST.
	Method string `json:"method,omitempty"`

	// Query contains the definitions for query parameters.
	// This schema MUST be of type object and have a properties key.
	Query interface{} `json:"query,omitempty"`

	// Headers contains the definitions of HTTP headers to use.
	// This schema MUST be of type object and have a properties key.
	Headers interface{} `json:"headers,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
