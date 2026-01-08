package solace

// OperationBinding represents the Solace Operation Binding object.
//
// This object contains information about the operation representation in Solace.
// +binding
type OperationBinding struct {
	// Destinations is a list of destinations used in the operation.
	Destinations []*Destination `json:"destinations,omitempty"`

	// TimeToLive is the time-to-live for messages in milliseconds.
	// Can be a Schema Object or Reference Object.
	TimeToLive interface{} `json:"timeToLive,omitempty"`

	// Priority is the message priority (0-255).
	// Can be a Schema Object or Reference Object.
	Priority interface{} `json:"priority,omitempty"`

	// DMQEligible indicates whether the message is eligible for the Dead Message Queue.
	DMQEligible bool `json:"dmqEligible,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}

// Destination represents a Solace destination configuration.
type Destination struct {
	// DestinationType is the type of destination ("queue" or "topic").
	DestinationType string `json:"destinationType,omitempty"`

	// DeliveryMode is the delivery mode ("direct" or "persistent").
	DeliveryMode string `json:"deliveryMode,omitempty"`

	// Queue contains queue-specific configuration.
	Queue *QueueDestination `json:"queue,omitempty"`

	// Topic contains topic-specific configuration.
	Topic *TopicDestination `json:"topic,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}

// QueueDestination represents queue-specific destination configuration.
type QueueDestination struct {
	// Name is the name of the queue.
	Name string `json:"name,omitempty"`

	// TopicSubscriptions is a list of topics that the queue subscribes to.
	TopicSubscriptions []string `json:"topicSubscriptions,omitempty"`

	// AccessType defines how clients access the queue ("exclusive" or "nonexclusive").
	AccessType string `json:"accessType,omitempty"`

	// MaxMsgSpoolSize is the maximum message spool size in MB.
	MaxMsgSpoolSize int `json:"maxMsgSpoolSize,omitempty"`

	// MaxTtl is the maximum TTL to apply to messages in seconds.
	MaxTtl int `json:"maxTtl,omitempty"`
}

// TopicDestination represents topic-specific destination configuration.
type TopicDestination struct {
	// TopicSubscriptions is a list of topic subscriptions.
	TopicSubscriptions []string `json:"topicSubscriptions,omitempty"`
}
