package ibmmq

// MessageBinding represents the IBM MQ Message Binding object.
//
// This object contains information about the message representation in IBM MQ.
// +binding
type MessageBinding struct {
	// Type is the type of the message ("string", "jms", or "binary").
	Type string `json:"type,omitempty"`

	// Headers contains the IBM MQ message headers for binary type messages.
	Headers string `json:"headers,omitempty"`

	// Description provides documentation on the message format.
	Description string `json:"description,omitempty"`

	// Expiry is the message time-to-live in milliseconds.
	Expiry int `json:"expiry,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
