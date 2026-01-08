package ibmmq

// OperationBinding represents the IBM MQ Operation Binding object.
//
// This object MUST NOT contain any properties. Its name is reserved for future use.
// +binding
type OperationBinding struct {
	BindingVersion string `json:"bindingVersion,omitempty"`
}
