package kafka

import (
	"encoding/json"
	"sigs.k8s.io/yaml"
)

// ChannelBinding represents the Kafka Channel Binding object.
// This object contains information about the channel representation in Kafka.
// +binding
// +binding:marshal:no-gen
type ChannelBinding struct {
	Topic              string              `json:"topic,omitempty"`
	Partitions         int                 `json:"partitions,omitempty"`
	Replicas           int                 `json:"replicas,omitempty"`
	TopicConfiguration *TopicConfiguration `json:"topicConfiguration,omitempty"`
	BindingVersion     string              `json:"bindingVersion,omitempty"`
}

// TopicConfiguration represents Kafka topic configuration properties.
type TopicConfiguration struct {
	CleanupPolicy                     []string               `json:"cleanup.policy,omitempty"`
	RetentionMs                       int64                  `json:"retention.ms,omitempty"`
	RetentionBytes                    int64                  `json:"retention.bytes,omitempty"`
	DeleteRetentionMs                 int64                  `json:"delete.retention.ms,omitempty"`
	MaxMessageBytes                   int                    `json:"max.message.bytes,omitempty"`
	ConfluentKeySchemaValidation      bool                   `json:"confluent.key.schema.validation,omitempty"`
	ConfluentKeySubjectNameStrategy   string                 `json:"confluent.key.subject.name.strategy,omitempty"`
	ConfluentValueSchemaValidation    bool                   `json:"confluent.value.schema.validation,omitempty"`
	ConfluentValueSubjectNameStrategy string                 `json:"confluent.value.subject.name.strategy,omitempty"`
	AdditionalProperties              map[string]interface{} `json:"-"`
}

// MarshalYAML is a custom marshaller that converts ChannelBinding to YAML
func (c ChannelBinding) MarshalYAML() (interface{}, error) {
	type Alias ChannelBinding
	if c.TopicConfiguration == nil {
		return struct{ Alias }{Alias(c)}, nil
	}

	// First marshal the ChannelBinding without TopicConfiguration
	temp := c
	temp.TopicConfiguration = nil
	m := make(map[string]interface{})
	bytes, err := json.Marshal(struct{ Alias }{Alias(temp)})
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &m); err != nil {
		return nil, err
	}

	// Then handle TopicConfiguration separately
	topicConfig := make(map[string]interface{})
	bytes, err = json.Marshal(c.TopicConfiguration)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &topicConfig); err != nil {
		return nil, err
	}

	// Add additional properties
	for k, v := range c.TopicConfiguration.AdditionalProperties {
		topicConfig[k] = v
	}
	m["topicConfiguration"] = topicConfig

	return m, nil
}

// UnmarshalYAML is a custom unmarshaler that converts YAML to ChannelBinding
func (c *ChannelBinding) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var temp interface{}
	if err := unmarshal(&temp); err != nil {
		return err
	}
	bytes, err := yaml.Marshal(temp)
	if err != nil {
		return err
	}

	type Alias ChannelBinding
	aux := struct{ *Alias }{Alias: (*Alias)(c)}
	if err := json.Unmarshal(bytes, &aux); err != nil {
		return err
	}

	// Handle additional properties in TopicConfiguration
	if c.TopicConfiguration != nil {
		var m map[string]interface{}
		if err := yaml.Unmarshal(bytes, &m); err != nil {
			return err
		}
		if topicConfig, ok := m["topicConfiguration"].(map[string]interface{}); ok {
			c.TopicConfiguration.AdditionalProperties = make(map[string]interface{})
			for k, v := range topicConfig {
				if k != "cleanup.policy" &&
					k != "retention.ms" &&
					k != "retention.bytes" &&
					k != "delete.retention.ms" &&
					k != "max.message.bytes" &&
					k != "confluent.key.schema.validation" &&
					k != "confluent.key.subject.name.strategy" &&
					k != "confluent.value.schema.validation" &&
					k != "confluent.value.subject.name.strategy" {
					c.TopicConfiguration.AdditionalProperties[k] = v
				}
			}
		}
	}
	return nil
}

// MarshalJSON is a custom marshaller that converts ChannelBinding to JSON
func (c ChannelBinding) MarshalJSON() ([]byte, error) {
	type Alias ChannelBinding
	if c.TopicConfiguration == nil {
		return json.Marshal(struct{ Alias }{Alias(c)})
	}

	// First marshal the ChannelBinding without TopicConfiguration
	temp := c
	temp.TopicConfiguration = nil
	m := make(map[string]interface{})
	bytes, err := json.Marshal(struct{ Alias }{Alias(temp)})
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &m); err != nil {
		return nil, err
	}

	// Then handle TopicConfiguration separately
	topicConfig := make(map[string]interface{})
	bytes, err = json.Marshal(c.TopicConfiguration)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bytes, &topicConfig); err != nil {
		return nil, err
	}

	// Add additional properties
	for k, v := range c.TopicConfiguration.AdditionalProperties {
		topicConfig[k] = v
	}
	m["topicConfiguration"] = topicConfig

	return json.Marshal(m)
}

// UnmarshalJSON is a custom unmarshaler that converts JSON to ChannelBinding
func (c *ChannelBinding) UnmarshalJSON(data []byte) error {
	type Alias ChannelBinding
	aux := struct{ *Alias }{Alias: (*Alias)(c)}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Handle additional properties in TopicConfiguration
	if c.TopicConfiguration != nil {
		var m map[string]interface{}
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}
		if topicConfig, ok := m["topicConfiguration"].(map[string]interface{}); ok {
			c.TopicConfiguration.AdditionalProperties = make(map[string]interface{})
			for k, v := range topicConfig {
				if k != "cleanup.policy" &&
					k != "retention.ms" &&
					k != "retention.bytes" &&
					k != "delete.retention.ms" &&
					k != "max.message.bytes" &&
					k != "confluent.key.schema.validation" &&
					k != "confluent.key.subject.name.strategy" &&
					k != "confluent.value.schema.validation" &&
					k != "confluent.value.subject.name.strategy" {
					c.TopicConfiguration.AdditionalProperties[k] = v
				}
			}
		}
	}
	return nil
}
