package ibmmq

// ServerBinding represents the IBM MQ Server Binding object.
//
// This object contains information about the server representation in IBM MQ.
// +binding
type ServerBinding struct {
	// GroupID is used to group IBM MQ endpoints (for HA clusters).
	GroupID string `json:"groupId,omitempty"`

	// CcdtQueueManagerName specifies the queue manager name when using a CCDT file.
	CcdtQueueManagerName string `json:"ccdtQueueManagerName,omitempty"`

	// CipherSpec is the recommended cipher specification for TLS connections.
	CipherSpec string `json:"cipherSpec,omitempty"`

	// MultiEndpointServer enables multi-endpoint client mode for workload balancing.
	MultiEndpointServer bool `json:"multiEndpointServer,omitempty"`

	// HeartBeatInterval is the interval in seconds for inactivity heartbeats.
	HeartBeatInterval int `json:"heartBeatInterval,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
