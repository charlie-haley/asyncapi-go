package nats

// ServerBinding represents the NATS Server Binding object.
//
// This object MUST NOT contain any properties. Its name is reserved for future use.
// +binding
type ServerBinding struct {
	BindingVersion string `json:"bindingVersion,omitempty"`
}
