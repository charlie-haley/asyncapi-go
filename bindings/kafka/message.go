package kafka

// MessageBinding represents the Kafka Message Binding object.
//
// This object contains information about the message representation in Kafka.
// +binding
type MessageBinding struct {
	Key                     interface{} `json:"key,omitempty"`
	SchemaIDLocation        string      `json:"schemaIdLocation,omitempty"`
	SchemaIDPayloadEncoding string      `json:"schemaIdPayloadEncoding,omitempty"`
	SchemaLookupStrategy    string      `json:"schemaLookupStrategy,omitempty"`
	BindingVersion          string      `json:"bindingVersion,omitempty"`
}
