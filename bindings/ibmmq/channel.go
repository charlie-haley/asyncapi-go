package ibmmq

// ChannelBinding represents the IBM MQ Channel Binding object.
//
// This object contains information about the channel representation in IBM MQ.
// +binding
type ChannelBinding struct {
	// DestinationType specifies whether the destination is a "topic" or "queue".
	DestinationType string `json:"destinationType,omitempty"`

	// Queue contains queue-specific properties.
	Queue *QueueProperties `json:"queue,omitempty"`

	// Topic contains topic-specific properties.
	Topic *TopicProperties `json:"topic,omitempty"`

	// MaxMsgLength is the maximum message length in bytes.
	MaxMsgLength int `json:"maxMsgLength,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}

// QueueProperties represents IBM MQ queue properties.
type QueueProperties struct {
	// ObjectName is the IBM MQ queue name.
	ObjectName string `json:"objectName"`

	// IsPartitioned indicates whether this is a cluster queue.
	IsPartitioned bool `json:"isPartitioned,omitempty"`

	// Exclusive recommends whether exclusive access is required.
	Exclusive bool `json:"exclusive,omitempty"`
}

// TopicProperties represents IBM MQ topic properties.
type TopicProperties struct {
	// String is the value of the IBM MQ topic string.
	String string `json:"string,omitempty"`

	// ObjectName is the name of the IBM MQ topic object.
	ObjectName string `json:"objectName,omitempty"`

	// DurablePermitted indicates whether durable subscriptions are permitted.
	DurablePermitted bool `json:"durablePermitted,omitempty"`

	// LastMsgRetained indicates whether retained publication is available.
	LastMsgRetained bool `json:"lastMsgRetained,omitempty"`
}
