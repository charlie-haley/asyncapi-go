package pulsar

// ServerBinding represents the Pulsar Server Binding object.
//
// This object contains information about the server representation in Pulsar.
// +binding
type ServerBinding struct {
	// Tenant is the pulsar tenant. If omitted, "public" MUST be assumed.
	Tenant string `json:"tenant,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
