package asyncapi

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/charlie-haley/asyncapi-go/asyncapi2"
	"github.com/charlie-haley/asyncapi-go/bindings/amqp"
	"github.com/charlie-haley/asyncapi-go/bindings/kafka"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestParseNestedRefs tests parsing of nested message and schema references
func TestParseNestedRefs(t *testing.T) {
	// Create temporary test directory
	tmpDir := t.TempDir()

	// Create schema file
	schemaJSON := `{
		"type": "object",
		"properties": {
			"eventId": {
				"type": "string",
				"format": "uuid"
			},
			"timestamp": {
				"type": "string",
				"format": "date-time"
			}
		},
		"required": ["eventId", "timestamp"]
	}`
	err := os.WriteFile(filepath.Join(tmpDir, "event-schema.json"), []byte(schemaJSON), 0644)
	require.NoError(t, err)

	// Create message file
	messageJSON := `{
		"name": "EventMessage",
		"contentType": "application/json",
		"payload": {
			"$ref": "./event-schema.json"
		}
	}`
	err = os.WriteFile(filepath.Join(tmpDir, "event-message.json"), []byte(messageJSON), 0644)
	require.NoError(t, err)

	// Create main AsyncAPI document
	asyncapiJSON := `{
		"asyncapi": "2.6.0",
		"info": {
			"title": "Test API",
			"version": "1.0.0"
		},
		"channels": {
			"user-signedup": {
				"publish": {
					"message": {
						"$ref": "./event-message.json"
					}
				}
			}
		}
	}`
	err = os.WriteFile(filepath.Join(tmpDir, "asyncapi.json"), []byte(asyncapiJSON), 0644)
	require.NoError(t, err)

	// Change to temp directory for test
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	err = os.Chdir(tmpDir)
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	// Parse the document
	data, err := os.ReadFile("asyncapi.json")
	require.NoError(t, err)

	doc, err := ParseFromJSON(data)
	require.NoError(t, err)
	require.NotNil(t, doc)

	// Convert to v2 document to access fields
	v2Doc, ok := doc.(*asyncapi2.Document)
	require.True(t, ok, "document should be v2")

	// Verify message and schema were properly resolved
	channel, ok := v2Doc.Channels["user-signedup"]
	require.True(t, ok, "channel should exist")
	require.NotNil(t, channel)

	require.NotNil(t, channel.Publish)
	require.NotNil(t, channel.Publish.Message)

	// Parse payload as Schema
	payload, ok := channel.Publish.Message.Payload.(map[string]interface{})
	require.True(t, ok, "payload should be a map")

	// Verify schema type
	assert.Equal(t, "object", payload["type"])

	// Check properties
	props, ok := payload["properties"].(map[string]interface{})
	require.True(t, ok, "should have properties")

	// Check eventId field
	eventId, ok := props["eventId"].(map[string]interface{})
	require.True(t, ok, "should have eventId")
	assert.Equal(t, "string", eventId["type"])
	assert.Equal(t, "uuid", eventId["format"])

	// Check timestamp field
	timestamp, ok := props["timestamp"].(map[string]interface{})
	require.True(t, ok, "should have timestamp")
	assert.Equal(t, "string", timestamp["type"])
	assert.Equal(t, "date-time", timestamp["format"])
}

// TestParseWithLocalRefs tests parsing of local references within the same document
func TestParseWithLocalRefs(t *testing.T) {
	asyncapiJSON := `{
		"asyncapi": "2.6.0",
		"info": {
			"title": "Test API",
			"version": "1.0.0"
		},
		"components": {
			"schemas": {
				"Event": {
					"type": "object",
					"properties": {
						"eventId": {
							"type": "string",
							"format": "uuid"
						}
					}
				}
			},
			"messages": {
				"EventMessage": {
					"payload": {
						"$ref": "#/components/schemas/Event"
					}
				}
			}
		},
		"channels": {
			"user-signedup": {
				"publish": {
					"message": {
						"$ref": "#/components/messages/EventMessage"
					}
				}
			}
		}
	}`

	doc, err := ParseFromJSON([]byte(asyncapiJSON))
	require.NoError(t, err)
	require.NotNil(t, doc)

	// Convert to v2 document to access fields
	v2Doc, ok := doc.(*asyncapi2.Document)
	require.True(t, ok, "document should be v2")

	// Verify message and schema were properly resolved
	channel, ok := v2Doc.Channels["user-signedup"]
	require.True(t, ok, "channel should exist")
	require.NotNil(t, channel)

	require.NotNil(t, channel.Publish)
	require.NotNil(t, channel.Publish.Message)

	// Parse payload as Schema
	payload, ok := channel.Publish.Message.Payload.(map[string]interface{})
	require.True(t, ok, "payload should be a map")

	// Verify schema type
	assert.Equal(t, "object", payload["type"])

	// Check properties
	props, ok := payload["properties"].(map[string]interface{})
	require.True(t, ok, "should have properties")

	// Check eventId field
	eventId, ok := props["eventId"].(map[string]interface{})
	require.True(t, ok, "should have eventId")
	assert.Equal(t, "string", eventId["type"])
	assert.Equal(t, "uuid", eventId["format"])
}

// Test individual specs
func TestParseV2Specs(t *testing.T) {
	tests := []struct {
		version     string
		spec        string
		expectError bool
	}{
		{
			version: "2.0.0",
			spec: `{
				"asyncapi": "2.0.0",
				"info": {
					"title": "Account Service",
					"version": "1.0.0"
				},
				"channels": {
					"user/signedup": {}
				}
			}`,
			expectError: false,
		},
		{
			version: "2.1.0",
			spec: `{
				"asyncapi": "2.1.0",
				"info": {
					"title": "Account Service",
					"version": "1.0.0"
				},
				"channels": {
					"user/signedup": {}
				}
			}`,
			expectError: false,
		},
		{
			version: "2.4.0",
			spec: `{
				"asyncapi": "2.4.0",
				"info": {
					"title": "Account Service",
					"version": "1.0.0"
				},
				"channels": {
					"user/signedup": {}
				}
			}`,
			expectError: false,
		},
		{
			version: "2.6.0",
			spec: `{
				"asyncapi": "2.6.0",
				"info": {
					"title": "Account Service",
					"version": "1.0.0"
				},
				"channels": {
					"user/signedup": {}
				}
			}`,
			expectError: false,
		},
		{
			version:     "3.0.0",
			spec:        `{"asyncapi": "3.0.0", "info": {"title": "Invalid Service", "version": "1.0.0"}, "channels": {}}`,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.version, func(t *testing.T) {
			doc, err := ParseFromJSON([]byte(tt.spec))
			if tt.expectError {
				assert.Error(t, err, "Expected error parsing spec")
				assert.Nil(t, doc, "Expected nil document on error")
			} else {
				assert.NoError(t, err, "Unexpected error parsing spec")
				require.NotNil(t, doc, "Expected non-nil document")
				assert.Equal(t, tt.version, doc.GetVersion(), "Expected AsyncAPI version to match")
			}
		})
	}
}

// Test parsing YAML documents
func TestParseFromYAML(t *testing.T) {
	yamlDoc := `
asyncapi: 2.6.0
info:
  title: YAML Account Service
  version: 1.0.0
channels:
  user/signedup: {}
`
	doc, err := ParseFromYAML([]byte(yamlDoc))
	assert.NoError(t, err, "Unexpected error parsing YAML")
	require.NotNil(t, doc, "Expected non-nil document")
	assert.Equal(t, "2.6.0", doc.GetVersion(), "Expected AsyncAPI version to match")
}

// Test parsing documents from testdata folder
func TestParseFromTestData(t *testing.T) {
	err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		t.Run(info.Name(), func(t *testing.T) {
			data, err := os.ReadFile(path)
			assert.NoError(t, err, "Error reading test file")

			doc, err := Parse(data)
			if strings.Contains(info.Name(), "invalid") {
				assert.Error(t, err, "Expected error parsing invalid spec")
				assert.Nil(t, doc, "Expected nil document on error")
			} else {
				assert.NoError(t, err, "Unexpected error parsing spec")
				require.NotNil(t, doc, "Expected non-nil document")
			}
		})

		return nil
	})

	assert.NoError(t, err, "Error walking testdata directory")
}

// Test isYAML function
func TestIsYAML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid YAML", "asyncapi: 2.0.0\ninfo:\n  title: Test", true},
		{"Valid JSON", `{"asyncapi": "2.0.0", "info": {"title": "Test"}}`, false},
		{"Empty String", "", true},
		{"Just Spaces", "   ", true},
		{"Invalid YAML", ":", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isYAML([]byte(tt.input)), "isYAML result mismatch")
		})
	}
}

// Test ParseBindings - AMQP
func TestParseBindings_AMQP(t *testing.T) {
	rawBindings := map[string]interface{}{
		"amqp": map[string]interface{}{
			"is": "routingKey",
			"exchange": map[string]interface{}{
				"name":       "myExchange",
				"type":       "topic",
				"durable":    true,
				"autoDelete": false,
				"vhost":      "/",
			},
			"bindingVersion": "0.3.0",
		},
	}

	expected := &amqp.ChannelBinding{
		Is: "routingKey",
		Exchange: &amqp.Exchange{
			Name:       "myExchange",
			Type:       "topic",
			Durable:    true,
			AutoDelete: false,
			VHost:      "/",
		},
		BindingVersion: "0.3.0",
	}

	binding, err := ParseBindings[amqp.ChannelBinding](rawBindings, "amqp")
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(expected, binding))
}

// Test ParseBindings - Kafka
func TestParseBindings_Kafka(t *testing.T) {
	rawBindings := map[string]interface{}{
		"kafka": map[string]interface{}{
			"topic":      "my-specific-topic-name",
			"partitions": 20,
			"replicas":   3,
			"topicConfiguration": map[string]interface{}{
				"cleanup.policy":      []string{"delete", "compact"},
				"retention.ms":        604800000,
				"retention.bytes":     1000000000,
				"delete.retention.ms": 86400000,
				"max.message.bytes":   1048588,
			},
			"bindingVersion": "0.5.0",
		},
	}

	expected := &kafka.ChannelBinding{
		Topic:      "my-specific-topic-name",
		Partitions: 20,
		Replicas:   3,
		TopicConfiguration: &kafka.TopicConfiguration{
			CleanupPolicy:     []string{"delete", "compact"},
			RetentionMs:       604800000,
			RetentionBytes:    1000000000,
			DeleteRetentionMs: 86400000,
			MaxMessageBytes:   1048588,
		},
		BindingVersion: "0.5.0",
	}

	binding, err := ParseBindings[kafka.ChannelBinding](rawBindings, "kafka")
	assert.NoError(t, err)

	// Handle the case where AdditionalProperties is nil in expected but an empty map in binding
	if expected.TopicConfiguration.AdditionalProperties == nil {
		expected.TopicConfiguration.AdditionalProperties = make(map[string]interface{})
	}

	assert.True(t, reflect.DeepEqual(expected, binding), "Expected: %+v, Actual: %+v", expected, binding)
}

// Test ParseBindings - Not Found
func TestParseBindings_NotFound(t *testing.T) {
	rawBindings := map[string]interface{}{
		"amqp": map[string]interface{}{
			"is":    "routingKey",
			"vhost": "/",
		},
	}

	_, err := ParseBindings[amqp.ChannelBinding](rawBindings, "http") // http doesn't exist
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "binding type http not found")
}

// Test ParseBindings - Invalid
func TestParseBindings_Invalid(t *testing.T) {
	rawBindings := map[string]interface{}{
		"amqp": "invalid", // invalid, not a map
	}

	_, err := ParseBindings[amqp.ChannelBinding](rawBindings, "amqp")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal amqp binding")
}

// Test ParseFromJSON - basic error check
func TestParseFromJSON_Error(t *testing.T) {
	_, err := ParseFromJSON([]byte("invalid json"))
	assert.Error(t, err)
}

// Test ParseFromYAML - basic error check
func TestParseFromYAML_Error(t *testing.T) {
	_, err := ParseFromYAML([]byte(": not yaml"))
	assert.Error(t, err)
}

// Test Parse - test both JSON and YAML
func TestParse(t *testing.T) {
	jsonDoc := `{"asyncapi": "2.6.0", "info": {"title": "Test Service", "version": "1.0.0"}, "channels": {"test": {}}}`
	yamlDoc := "asyncapi: 2.6.0\ninfo:\n  title: Test Service\n  version: 1.0.0\nchannels:\n  test: {}"

	jsonParsed, err := Parse([]byte(jsonDoc))
	assert.NoError(t, err)
	require.NotNil(t, jsonParsed)

	yamlParsed, err := Parse([]byte(yamlDoc))
	assert.NoError(t, err)
	require.NotNil(t, yamlParsed)
}
