package googlepubsub

// MessageBinding represents the Google Cloud Pub/Sub Message Binding object.
// +binding
type MessageBinding struct {
	Attributes     map[string]string `json:"attributes,omitempty"`
	OrderingKey    string            `json:"orderingKey,omitempty"`
	Schema         *Schema           `json:"schema,omitempty"`
	BindingVersion string            `json:"bindingVersion,omitempty"`
}

// Schema represents a Google Cloud Pub/Sub schema reference.
type Schema struct {
	Name string `json:"name,omitempty"`
}
