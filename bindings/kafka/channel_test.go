package kafka

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestChannelBinding_TopicConfigurationAdditionalProperties(t *testing.T) {
	jsonString := `{
        "topic": "test-topic",
        "topicConfiguration": {
            "cleanup.policy": ["delete"],
            "custom.property": "custom-value",
            "another.property": 123,
            "nested.property": {"key": "value"}
        }
    }`

	cb := NewChannelBinding()
	err := cb.UnmarshalJSON([]byte(jsonString))
	assert.NoError(t, err)
	assert.Equal(t, "custom-value", cb.TopicConfiguration.AdditionalProperties["custom.property"])
	assert.Equal(t, float64(123), cb.TopicConfiguration.AdditionalProperties["another.property"])
	assert.Equal(t, map[string]interface{}{"key": "value"}, cb.TopicConfiguration.AdditionalProperties["nested.property"])

	marshaledJSON, err := cb.MarshalJSON()
	assert.NoError(t, err)

	unmarshaled := NewChannelBinding()
	err = unmarshaled.UnmarshalJSON(marshaledJSON)
	assert.NoError(t, err)

	assert.Equal(t, "custom-value", unmarshaled.TopicConfiguration.AdditionalProperties["custom.property"])
	assert.Equal(t, float64(123), unmarshaled.TopicConfiguration.AdditionalProperties["another.property"])
	assert.Equal(t, map[string]interface{}{"key": "value"}, unmarshaled.TopicConfiguration.AdditionalProperties["nested.property"])
}

func TestChannelBinding_TopicConfigurationYAMLAdditionalProperties(t *testing.T) {
	yamlString := []byte(`
topic: test-topic
topicConfiguration:
  cleanup.policy:
    - delete
  custom.property: custom-value
  another.property: 123
  nested.property:
    key: value
`)

	cb := NewChannelBinding()
	err := yaml.Unmarshal(yamlString, &cb)
	assert.NoError(t, err)
	assert.Equal(t, "custom-value", cb.TopicConfiguration.AdditionalProperties["custom.property"])
	assert.Equal(t, float64(123), cb.TopicConfiguration.AdditionalProperties["another.property"])
	assert.Equal(t, map[string]interface{}{"key": "value"}, cb.TopicConfiguration.AdditionalProperties["nested.property"])

	marshaledYAML, err := yaml.Marshal(cb)
	assert.NoError(t, err)

	unmarshaled := NewChannelBinding()
	err = yaml.Unmarshal(marshaledYAML, unmarshaled)
	assert.NoError(t, err)

	assert.Equal(t, "custom-value", unmarshaled.TopicConfiguration.AdditionalProperties["custom.property"])
	assert.Equal(t, float64(123), unmarshaled.TopicConfiguration.AdditionalProperties["another.property"])
	assert.Equal(t, map[string]interface{}{"key": "value"}, unmarshaled.TopicConfiguration.AdditionalProperties["nested.property"])
}

func TestChannelBinding_BuildObject(t *testing.T) {
	cb := NewChannelBinding().
		WithTopic("my-specific-topic-name").
		WithPartitions(20).
		WithReplicas(3).
		WithTopicConfiguration(NewTopicConfiguration().
			WithCleanupPolicy([]string{"delete", "compact"}).
			WithRetentionMs(604800000).
			WithRetentionBytes(1000000000).
			WithDeleteRetentionMs(86400000).
			WithMaxMessageBytes(1048588),
		)

	assert.Equal(t, "my-specific-topic-name", cb.Topic)
	assert.Equal(t, 20, cb.Partitions)
	assert.Equal(t, 3, cb.Replicas)
	assert.NotNil(t, cb.TopicConfiguration)
	assert.Equal(t, []string{"delete", "compact"}, cb.TopicConfiguration.CleanupPolicy)
	assert.Equal(t, int64(604800000), cb.TopicConfiguration.RetentionMs)
	assert.Equal(t, int64(1000000000), cb.TopicConfiguration.RetentionBytes)
	assert.Equal(t, int64(86400000), cb.TopicConfiguration.DeleteRetentionMs)
	assert.Equal(t, 1048588, cb.TopicConfiguration.MaxMessageBytes)
}

func TestChannelBinding_MarshalYAML(t *testing.T) {
	cb := NewChannelBinding().
		WithTopic("my-specific-topic-name").
		WithPartitions(20).
		WithReplicas(3).
		WithTopicConfiguration(NewTopicConfiguration().
			WithCleanupPolicy([]string{"delete", "compact"}).
			WithRetentionMs(604800000),
		)

	expectedYAML := `partitions: 20
replicas: 3
topic: my-specific-topic-name
topicConfiguration:
  cleanup.policy:
  - delete
  - compact
  retention.ms: 604800000
`
	marshaledYAML, err := yaml.Marshal(cb)
	assert.NoError(t, err)
	assert.Equal(t, expectedYAML, string(marshaledYAML))
}

func TestChannelBinding_UnmarshalYAML(t *testing.T) {
	yamlString := `
topic: test-topic
partitions: 10
replicas: 2
topicConfiguration:
  cleanup.policy:
    - delete
  retention.ms: 300000
  retention.bytes: 500000000
  delete.retention.ms: 43200000
  max.message.bytes: 524288
  confluent.key.schema.validation: true
  confluent.key.subject.name.strategy: TopicNameStrategy
  confluent.value.schema.validation: true
  confluent.value.subject.name.strategy: TopicNameStrategy
`

	var cb ChannelBinding
	err := yaml.Unmarshal([]byte(yamlString), &cb)
	assert.NoError(t, err)

	assert.Equal(t, "test-topic", cb.Topic)
	assert.Equal(t, 10, cb.Partitions)
	assert.Equal(t, 2, cb.Replicas)
	assert.NotNil(t, cb.TopicConfiguration)
	assert.Equal(t, []string{"delete"}, cb.TopicConfiguration.CleanupPolicy)
	assert.Equal(t, int64(300000), cb.TopicConfiguration.RetentionMs)
	assert.Equal(t, int64(500000000), cb.TopicConfiguration.RetentionBytes)
	assert.Equal(t, int64(43200000), cb.TopicConfiguration.DeleteRetentionMs)
	assert.Equal(t, 524288, cb.TopicConfiguration.MaxMessageBytes)
	assert.True(t, cb.TopicConfiguration.ConfluentKeySchemaValidation)
	assert.Equal(t, "TopicNameStrategy", cb.TopicConfiguration.ConfluentKeySubjectNameStrategy)
	assert.True(t, cb.TopicConfiguration.ConfluentValueSchemaValidation)
	assert.Equal(t, "TopicNameStrategy", cb.TopicConfiguration.ConfluentValueSubjectNameStrategy)
}

func TestChannelBinding_MarshalJSON(t *testing.T) {
	cb := NewChannelBinding().
		WithTopic("my-specific-topic-name").
		WithPartitions(20).
		WithReplicas(3).
		WithTopicConfiguration(NewTopicConfiguration().
			WithCleanupPolicy([]string{"delete", "compact"}).
			WithRetentionMs(604800000),
		)

	expectedJSON := `{"partitions":20,"replicas":3,"topic":"my-specific-topic-name","topicConfiguration":{"cleanup.policy":["delete","compact"],"retention.ms":604800000}}`

	marshaledJSON, err := json.Marshal(cb)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(marshaledJSON))
}

func TestChannelBinding_UnmarshalJSON(t *testing.T) {
	jsonString := `{"topic":"test-topic","partitions":10,"replicas":2,"topicConfiguration":{"cleanup.policy":["delete"],"retention.ms":300000,"retention.bytes":500000000,"delete.retention.ms":43200000,"max.message.bytes":524288,"confluent.key.schema.validation":true,"confluent.key.subject.name.strategy":"TopicNameStrategy","confluent.value.schema.validation":true,"confluent.value.subject.name.strategy":"TopicNameStrategy"}}`

	var cb ChannelBinding
	err := json.Unmarshal([]byte(jsonString), &cb)
	assert.NoError(t, err)
	assert.Equal(t, "test-topic", cb.Topic)
	assert.Equal(t, 10, cb.Partitions)
	assert.Equal(t, 2, cb.Replicas)
	assert.NotNil(t, cb.TopicConfiguration)
	assert.Equal(t, []string{"delete"}, cb.TopicConfiguration.CleanupPolicy)
	assert.Equal(t, int64(300000), cb.TopicConfiguration.RetentionMs)
	assert.Equal(t, int64(500000000), cb.TopicConfiguration.RetentionBytes)
	assert.Equal(t, int64(43200000), cb.TopicConfiguration.DeleteRetentionMs)
	assert.Equal(t, 524288, cb.TopicConfiguration.MaxMessageBytes)
	assert.True(t, cb.TopicConfiguration.ConfluentKeySchemaValidation)
	assert.Equal(t, "TopicNameStrategy", cb.TopicConfiguration.ConfluentKeySubjectNameStrategy)
	assert.True(t, cb.TopicConfiguration.ConfluentValueSchemaValidation)
	assert.Equal(t, "TopicNameStrategy", cb.TopicConfiguration.ConfluentValueSubjectNameStrategy)
}
