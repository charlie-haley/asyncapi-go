package sns

// ChannelBinding represents the SNS Channel Binding object.
// This object contains information about the channel representation in SNS.
// +binding
type ChannelBinding struct {
	Name           string            `json:"name,omitempty"`
	Ordering       *Ordering         `json:"ordering,omitempty"`
	Policy         *Policy           `json:"policy,omitempty"`
	Tags           map[string]string `json:"tags,omitempty"`
	BindingVersion string            `json:"bindingVersion,omitempty"`
}

// Ordering represents SNS topic ordering configuration
type Ordering struct {
	Type                     string `json:"type"`
	ContentBasedDeduplication bool   `json:"contentBasedDeduplication,omitempty"`
}

// Policy represents an SNS topic access policy
type Policy struct {
	Statements []Statement `json:"statements"`
}

// Statement represents a single SNS policy statement
type Statement struct {
	Effect     string      `json:"effect"`
	Principal  interface{} `json:"principal"`
	Action     interface{} `json:"action"`
	Resource   interface{} `json:"resource,omitempty"`
	Condition  interface{} `json:"condition,omitempty"`
}