package anypointmq

// ChannelBinding represents the Anypoint MQ Channel Binding object.
//
// This object contains information about the channel representation in Anypoint MQ.
// +binding
type ChannelBinding struct {
	// Destination is the destination name. Defaults to the channel name.
	Destination string `json:"destination,omitempty"`

	// DestinationType is the type of destination.
	// MUST be one of "exchange", "queue", or "fifo-queue".
	// Defaults to "queue".
	DestinationType string `json:"destinationType,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
