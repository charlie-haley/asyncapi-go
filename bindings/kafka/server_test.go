package kafka

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestServerBinding_BuildObject(t *testing.T) {
	sb := NewServerBinding().
		WithSchemaRegistryURL("http://schema-registry:8081").
		WithSchemaRegistryVendor("confluent")

	assert.Equal(t, "http://schema-registry:8081", sb.SchemaRegistryURL)
	assert.Equal(t, "confluent", sb.SchemaRegistryVendor)
}

func TestServerBinding_MarshalYAML(t *testing.T) {
	sb := NewServerBinding().
		WithSchemaRegistryURL("http://schema-registry:8081").
		WithSchemaRegistryVendor("confluent")

	expectedYAML := `schemaRegistryUrl: http://schema-registry:8081
schemaRegistryVendor: confluent
`
	marshaledYAML, err := yaml.Marshal(sb)
	assert.NoError(t, err)
	assert.Equal(t, expectedYAML, string(marshaledYAML))
}

func TestServerBinding_UnmarshalYAML(t *testing.T) {
	yamlString := `
schemaRegistryUrl: http://schema-registry:8081
schemaRegistryVendor: confluent
`
	var sb ServerBinding
	err := yaml.Unmarshal([]byte(yamlString), &sb)
	assert.NoError(t, err)

	assert.Equal(t, "http://schema-registry:8081", sb.SchemaRegistryURL)
	assert.Equal(t, "confluent", sb.SchemaRegistryVendor)
}

func TestServerBinding_MarshalJSON(t *testing.T) {
	sb := NewServerBinding().
		WithSchemaRegistryURL("http://schema-registry:8081").
		WithSchemaRegistryVendor("confluent")

	expectedJSON := `{"schemaRegistryUrl":"http://schema-registry:8081","schemaRegistryVendor":"confluent"}`

	marshaledJSON, err := json.Marshal(sb)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(marshaledJSON))
}

func TestServerBinding_UnmarshalJSON(t *testing.T) {
	jsonString := `{
		"schemaRegistryUrl": "http://schema-registry:8081",
		"schemaRegistryVendor": "confluent"
	}`

	var sb ServerBinding
	err := json.Unmarshal([]byte(jsonString), &sb)
	assert.NoError(t, err)

	assert.Equal(t, "http://schema-registry:8081", sb.SchemaRegistryURL)
	assert.Equal(t, "confluent", sb.SchemaRegistryVendor)
}
