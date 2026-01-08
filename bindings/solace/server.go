package solace

// ServerBinding represents the Solace Server Binding object.
//
// This object contains information about the server representation in Solace.
// +binding
type ServerBinding struct {
	// MsgVpn is the Virtual Private Network name on the Solace broker.
	MsgVpn string `json:"msgVpn,omitempty"`

	// ClientName is a unique client name to use to register to the appliance.
	// Maximum 160 bytes of UTF-8 characters.
	ClientName string `json:"clientName,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
