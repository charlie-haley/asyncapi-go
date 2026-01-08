package mqtt

// ServerBinding represents the MQTT Server Binding object.
//
// This object contains information about the server representation in MQTT.
// +binding
type ServerBinding struct {
	// ClientID is the client identifier.
	ClientID string `json:"clientId,omitempty"`

	// CleanSession indicates whether to create a persistent connection.
	CleanSession bool `json:"cleanSession,omitempty"`

	// LastWill is the Last Will and Testament configuration.
	LastWill *LastWill `json:"lastWill,omitempty"`

	// KeepAlive is the interval in seconds of the longest period of time
	// the broker and client can endure without sending a message.
	KeepAlive int `json:"keepAlive,omitempty"`

	// SessionExpiryInterval is the session expiry interval in seconds.
	// Can be a Schema Object or Reference Object.
	SessionExpiryInterval interface{} `json:"sessionExpiryInterval,omitempty"`

	// MaximumPacketSize is the maximum packet size the client is willing to accept.
	// Can be a Schema Object or Reference Object.
	MaximumPacketSize interface{} `json:"maximumPacketSize,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}

// LastWill represents the Last Will and Testament configuration.
type LastWill struct {
	// Topic is the topic for the Last Will message.
	Topic string `json:"topic,omitempty"`

	// QoS is the quality of service level (0, 1, or 2).
	QoS int `json:"qos,omitempty"`

	// Message is the Last Will message content.
	Message string `json:"message,omitempty"`

	// Retain indicates whether the broker should retain the Last Will message.
	Retain bool `json:"retain,omitempty"`
}
