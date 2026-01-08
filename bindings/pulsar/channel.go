package pulsar

// ChannelBinding represents the Pulsar Channel Binding object.
//
// This object contains information about the channel representation in Pulsar.
// +binding
type ChannelBinding struct {
	// Namespace is the namespace the channel is associated with.
	Namespace string `json:"namespace"`

	// Persistence defines whether the topic is persistent or non-persistent.
	// MUST be either "persistent" or "non-persistent".
	Persistence string `json:"persistence"`

	// Compaction is the topic compaction threshold in Megabytes.
	Compaction int `json:"compaction,omitempty"`

	// GeoReplication is a list of clusters the topic is replicated to.
	GeoReplication []string `json:"geo-replication,omitempty"`

	// Retention defines the topic retention policy.
	Retention *RetentionDefinition `json:"retention,omitempty"`

	// TTL is the message time-to-live in seconds.
	TTL int `json:"ttl,omitempty"`

	// Deduplication indicates whether message deduplication is enabled.
	// When true, each message produced on Pulsar topics is persisted to disk only once.
	Deduplication bool `json:"deduplication,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}

// RetentionDefinition represents the topic retention policy.
type RetentionDefinition struct {
	// Time is the retention time in minutes.
	Time int `json:"time,omitempty"`

	// Size is the retention size in Megabytes.
	Size int `json:"size,omitempty"`
}
