package http

// ChannelBinding represents the HTTP Channel Binding object.
//
// This object MUST NOT contain any properties. Its name is reserved for future use.
// +binding
type ChannelBinding struct {
	BindingVersion string `json:"bindingVersion,omitempty"`
}
