package anypointmq

// MessageBinding represents the Anypoint MQ Message Binding object.
//
// This object contains information about the message representation in Anypoint MQ.
// +binding
type MessageBinding struct {
	// Headers is a schema defining Anypoint MQ protocol headers.
	// Includes messageId, messageGroupId, etc.
	Headers interface{} `json:"headers,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
