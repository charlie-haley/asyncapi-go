package googlepubsub

// ChannelBinding represents the Google Cloud Pub/Sub Channel Binding object.
// This object contains information about the channel representation in Google Cloud Pub/Sub.
// +binding
type ChannelBinding struct {
	Topic                    string                `json:"topic,omitempty"`
	MessageRetentionDuration string                `json:"messageRetentionDuration,omitempty"`
	MessageStoragePolicy     *MessageStoragePolicy `json:"messageStoragePolicy,omitempty"`
	SchemaSettings           *SchemaSettings       `json:"schemaSettings,omitempty"`
	Labels                   map[string]string     `json:"labels,omitempty"`
	BindingVersion           string                `json:"bindingVersion,omitempty"`
}

// MessageStoragePolicy represents the Google Cloud Pub/Sub message storage policy.
type MessageStoragePolicy struct {
	AllowedPersistenceRegions []string `json:"allowedPersistenceRegions,omitempty"`
}

// SchemaSettings represents the Google Cloud Pub/Sub schema settings.
type SchemaSettings struct {
	Encoding        string `json:"encoding,omitempty"`
	Name            string `json:"name,omitempty"`
	FirstRevisionID string `json:"firstRevisionId,omitempty"`
	LastRevisionID  string `json:"lastRevisionId,omitempty"`
}
