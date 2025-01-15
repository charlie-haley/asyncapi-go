package sns

// OperationBinding represents the SNS Operation Binding object.
// This object contains information about the operation representation in SNS.
// +binding
type OperationBinding struct {
	Topic          *Identifier      `json:"topic,omitempty"`
	Consumers      []Consumer       `json:"consumers,omitempty"`
	DeliveryPolicy *DeliveryPolicy `json:"deliveryPolicy,omitempty"`
	BindingVersion string          `json:"bindingVersion,omitempty"`
}

// Consumer represents an SNS topic subscriber configuration
type Consumer struct {
	Protocol           string          `json:"protocol"`
	Endpoint           Identifier      `json:"endpoint"`
	FilterPolicy       interface{}     `json:"filterPolicy,omitempty"`
	FilterPolicyScope  string          `json:"filterPolicyScope,omitempty"`
	RawMessageDelivery bool            `json:"rawMessageDelivery"`
	RedrivePolicy     *RedrivePolicy  `json:"redrivePolicy,omitempty"`
	DeliveryPolicy    *DeliveryPolicy `json:"deliveryPolicy,omitempty"`
	DisplayName       string          `json:"displayName,omitempty"`
}

// Identifier represents various ways to identify an SNS endpoint
type Identifier struct {
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	ARN   string `json:"arn,omitempty"`
	Name  string `json:"name,omitempty"`
}

// DeliveryPolicy represents SNS message delivery retry configuration
type DeliveryPolicy struct {
	MinDelayTarget       *int    `json:"minDelayTarget,omitempty"`
	MaxDelayTarget       *int    `json:"maxDelayTarget,omitempty"`
	NumRetries           *int    `json:"numRetries,omitempty"`
	NumNoDelayRetries    *int    `json:"numNoDelayRetries,omitempty"`
	NumMinDelayRetries   *int    `json:"numMinDelayRetries,omitempty"`
	NumMaxDelayRetries   *int    `json:"numMaxDelayRetries,omitempty"`
	BackoffFunction      string  `json:"backoffFunction,omitempty"`
	MaxReceivesPerSecond *int    `json:"maxReceivesPerSecond,omitempty"`
}

// RedrivePolicy represents SNS dead-letter queue configuration
type RedrivePolicy struct {
	DeadLetterQueue Identifier `json:"deadLetterQueue"`
	MaxReceiveCount *int       `json:"maxReceiveCount,omitempty"`
}