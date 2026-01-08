package jms

// ServerBinding represents the JMS Server Binding object.
//
// This object contains information about the server representation in JMS.
// +binding
type ServerBinding struct {
	// JmsConnectionFactory is the classname of the ConnectionFactory implementation.
	JmsConnectionFactory string `json:"jmsConnectionFactory"`

	// Properties is an array of additional properties for the JMS ConnectionFactory.
	Properties []Property `json:"properties,omitempty"`

	// ClientID is the client identifier for JMS connections.
	ClientID string `json:"clientID,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}

// Property represents a JMS property.
type Property struct {
	// Name is the property name.
	Name string `json:"name"`

	// Value is the property value.
	Value interface{} `json:"value"`
}
