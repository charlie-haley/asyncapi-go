package sqs

// OperationBinding represents the SQS Operation Binding object.
// This object contains information about the operation representation in SQS.
// +binding
type OperationBinding struct {
	Queues          []Queue `json:"queues,omitempty"`
	BindingVersion  string  `json:"bindingVersion,omitempty"`
}