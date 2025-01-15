package kafka

// OperationBinding represents the Kafka Operation Binding object.
//
// This object contains information about the operation representation in Kafka.
// +binding
type OperationBinding struct {
	GroupID        interface{} `json:"groupId,omitempty"`
	ClientID       interface{} `json:"clientId,omitempty"`
	BindingVersion string      `json:"bindingVersion,omitempty"`
}
