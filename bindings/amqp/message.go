package amqp

// MessageBinding represents the AMQP 0-9-1 Message Binding object.
//
// This object contains information about the message representation in AMQP.
// +binding
type MessageBinding struct {
	ContentEncoding string `json:"contentEncoding,omitempty"`
	MessageType     string `json:"messageType,omitempty"`
	BindingVersion  string `json:"bindingVersion,omitempty"`
}
