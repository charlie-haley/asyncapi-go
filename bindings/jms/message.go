package jms

// MessageBinding represents the JMS Message Binding object.
//
// This object contains information about the message representation in JMS.
// +binding
type MessageBinding struct {
	// Headers is a schema defining JMS protocol headers.
	// Includes JMSMessageID, JMSTimestamp, JMSDeliveryMode, JMSPriority, etc.
	Headers interface{} `json:"headers,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
