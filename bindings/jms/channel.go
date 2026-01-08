package jms

// ChannelBinding represents the JMS Channel Binding object.
//
// This object contains information about the channel representation in JMS.
// +binding
type ChannelBinding struct {
	// Destination is the JMS destination name. Defaults to the channel name.
	Destination string `json:"destination,omitempty"`

	// DestinationType is the type of destination ("queue" or "fifo-queue").
	// Defaults to "queue".
	DestinationType string `json:"destinationType,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
