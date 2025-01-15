package sqs

// ChannelBinding represents the SQS Channel Binding object.
// This object contains information about the channel representation in SQS.
// +binding
type ChannelBinding struct {
	Queue            *Queue   `json:"queue"`
	DeadLetterQueue  *Queue   `json:"deadLetterQueue,omitempty"`
	BindingVersion   string   `json:"bindingVersion,omitempty" default:"\"latest\""`
}

// Queue represents an SQS queue definition
type Queue struct {
	Name                   string         `json:"name"`
	FifoQueue             bool           `json:"fifoQueue"`
	DeduplicationScope    string         `json:"deduplicationScope,omitempty"`
	FifoThroughputLimit   string         `json:"fifoThroughputLimit,omitempty"`
	DeliveryDelay        int            `json:"deliveryDelay,omitempty"`
	VisibilityTimeout    int            `json:"visibilityTimeout,omitempty"`
	ReceiveMessageWaitTime int           `json:"receiveMessageWaitTime,omitempty"`
	MessageRetentionPeriod int           `json:"messageRetentionPeriod,omitempty"`
	RedrivePolicy        *RedrivePolicy `json:"redrivePolicy,omitempty"`
	Policy               *Policy        `json:"policy,omitempty"`
	Tags                 map[string]string `json:"tags,omitempty"`
}

// Identifier represents ways to identify an SQS queue
type Identifier struct {
	ARN  string `json:"arn,omitempty"`
	Name string `json:"name,omitempty"`
}

// Policy represents an SQS queue access policy
type Policy struct {
	Statements []Statement `json:"statements"`
}

// Statement represents a single SQS policy statement
type Statement struct {
	Effect     string      `json:"effect"`
	Principal  interface{} `json:"principal"`
	Action     interface{} `json:"action"`
	Resource   interface{} `json:"resource,omitempty"`
	Condition  interface{} `json:"condition,omitempty"`
}

// RedrivePolicy represents SQS dead-letter queue configuration
type RedrivePolicy struct {
	DeadLetterQueue Identifier `json:"deadLetterQueue"`
	MaxReceiveCount *int       `json:"maxReceiveCount,omitempty"`
}