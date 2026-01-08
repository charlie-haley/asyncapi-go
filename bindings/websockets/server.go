package websockets

// ServerBinding represents the WebSockets Server Binding object.
//
// This object MUST NOT contain any properties. Its name is reserved for future use.
// +binding
type ServerBinding struct {
	BindingVersion string `json:"bindingVersion,omitempty"`
}
