package websockets

// MessageBinding represents the WebSockets Message Binding object.
//
// This object MUST NOT contain any properties. Its name is reserved for future use.
// +binding
type MessageBinding struct {
	BindingVersion string `json:"bindingVersion,omitempty"`
}
