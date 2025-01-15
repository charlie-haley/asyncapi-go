package kafka

// ServerBinding represents the Kafka Server Binding object.
//
// This object contains information about the server representation in Kafka.
// +binding
type ServerBinding struct {
	SchemaRegistryURL    string `json:"schemaRegistryUrl,omitempty"`
	SchemaRegistryVendor string `json:"schemaRegistryVendor,omitempty"`
	BindingVersion       string `json:"bindingVersion,omitempty"`
}
