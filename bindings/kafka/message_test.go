package kafka

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestMessageBinding_BuildObject(t *testing.T) {
	mb := NewMessageBinding().
		WithKey("test-key").
		WithSchemaIDLocation("header").
		WithSchemaIDPayloadEncoding("confluent").
		WithSchemaLookupStrategy("TopicRecordNameStrategy")

	assert.Equal(t, "test-key", mb.Key)
	assert.Equal(t, "header", mb.SchemaIDLocation)
	assert.Equal(t, "confluent", mb.SchemaIDPayloadEncoding)
	assert.Equal(t, "TopicRecordNameStrategy", mb.SchemaLookupStrategy)
}

func TestMessageBinding_MarshalYAML(t *testing.T) {
	mb := NewMessageBinding().
		WithKey("test-key").
		WithSchemaIDLocation("header").
		WithSchemaIDPayloadEncoding("confluent").
		WithSchemaLookupStrategy("TopicRecordNameStrategy")

	expectedYAML := `key: test-key
schemaIdLocation: header
schemaIdPayloadEncoding: confluent
schemaLookupStrategy: TopicRecordNameStrategy
`
	marshaledYAML, err := yaml.Marshal(mb)
	assert.NoError(t, err)
	assert.Equal(t, expectedYAML, string(marshaledYAML))
}

func TestMessageBinding_UnmarshalYAML(t *testing.T) {
	yamlString := `
key: test-key
schemaIdLocation: header
schemaIdPayloadEncoding: confluent
schemaLookupStrategy: TopicRecordNameStrategy
`
	var mb MessageBinding
	err := yaml.Unmarshal([]byte(yamlString), &mb)
	assert.NoError(t, err)

	assert.Equal(t, "test-key", mb.Key)
	assert.Equal(t, "header", mb.SchemaIDLocation)
	assert.Equal(t, "confluent", mb.SchemaIDPayloadEncoding)
	assert.Equal(t, "TopicRecordNameStrategy", mb.SchemaLookupStrategy)
}

func TestMessageBinding_MarshalJSON(t *testing.T) {
	mb := NewMessageBinding().
		WithKey("test-key").
		WithSchemaIDLocation("header").
		WithSchemaIDPayloadEncoding("confluent").
		WithSchemaLookupStrategy("TopicRecordNameStrategy")

	expectedJSON := `{"key":"test-key","schemaIdLocation":"header","schemaIdPayloadEncoding":"confluent","schemaLookupStrategy":"TopicRecordNameStrategy"}`

	marshaledJSON, err := json.Marshal(mb)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(marshaledJSON))
}

func TestMessageBinding_UnmarshalJSON(t *testing.T) {
	jsonString := `{
		"key": "test-key",
		"schemaIdLocation": "header",
		"schemaIdPayloadEncoding": "confluent",
		"schemaLookupStrategy": "TopicRecordNameStrategy"
	}`

	var mb MessageBinding
	err := json.Unmarshal([]byte(jsonString), &mb)
	assert.NoError(t, err)

	assert.Equal(t, "test-key", mb.Key)
	assert.Equal(t, "header", mb.SchemaIDLocation)
	assert.Equal(t, "confluent", mb.SchemaIDPayloadEncoding)
	assert.Equal(t, "TopicRecordNameStrategy", mb.SchemaLookupStrategy)
}

func TestMessageBinding_NonStringKey(t *testing.T) {
	jsonString := `{
		"key": 123,
		"schemaIdLocation": "header"
	}`

	var mb MessageBinding
	err := json.Unmarshal([]byte(jsonString), &mb)
	assert.NoError(t, err)
	assert.Equal(t, float64(123), mb.Key)
}
