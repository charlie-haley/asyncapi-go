package nats

// OperationBinding represents the NATS Operation Binding object.
//
// This object contains information about the operation representation in NATS.
// +binding
type OperationBinding struct {
	// Queue specifies the name of the queue to use.
	// A queue group name allows clients to subscribe under a common name,
	// with messages being distributed among members of the group.
	// Maximum length is 255 characters.
	Queue string `json:"queue,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
