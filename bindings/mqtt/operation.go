package mqtt

// OperationBinding represents the MQTT Operation Binding object.
//
// This object contains information about the operation representation in MQTT.
// +binding
type OperationBinding struct {
	// QoS is the quality of service level for the message delivery.
	// Can be 0 (at most once), 1 (at least once), or 2 (exactly once).
	QoS int `json:"qos,omitempty"`

	// Retain indicates whether the broker should retain the message.
	Retain bool `json:"retain,omitempty"`

	// MessageExpiryInterval is the lifetime of the message in seconds.
	// Can be a Schema Object or Reference Object.
	MessageExpiryInterval interface{} `json:"messageExpiryInterval,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
