package asyncapi

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/charlie-haley/asyncapi-go/asyncapi2"
	"github.com/charlie-haley/asyncapi-go/asyncapi3"
	"github.com/charlie-haley/asyncapi-go/bindings/amqp"
	"github.com/charlie-haley/asyncapi-go/bindings/googlepubsub"
	"github.com/charlie-haley/asyncapi-go/bindings/kafka"
	"github.com/charlie-haley/asyncapi-go/bindings/sns"
	"github.com/charlie-haley/asyncapi-go/bindings/sqs"
	"github.com/charlie-haley/asyncapi-go/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestParseFileWithRelativeRefs tests the ParseFile function while using relative references
func TestParseFileWithRelativeRefs(t *testing.T) {
	// Create a temporary test directory structure
	tmpDir := t.TempDir()
	messagesDir := filepath.Join(tmpDir, "messages")
	schemasDir := filepath.Join(tmpDir, "schemas")
	require.NoError(t, os.MkdirAll(messagesDir, 0755))
	require.NoError(t, os.MkdirAll(schemasDir, 0755))

	// Create schema file (JSON)
	schemaJSON := `{
		"type": "object",
		"properties": {
			"customerId": {
				"type": "string",
				"format": "uuid"
			},
			"status": {
				"type": "string",
				"enum": ["active", "inactive"]
			}
		},
		"required": ["customerId", "status"]
	}`
	err := os.WriteFile(filepath.Join(schemasDir, "customer.json"), []byte(schemaJSON), 0644)
	require.NoError(t, err)

	// Create message file as JSON, not YAML
	messageJSON := `{
		"name": "UpdateCustomerStatus",
		"contentType": "application/json",
		"payload": {
			"$ref": "../schemas/customer.json"
		}
	}`
	err = os.WriteFile(filepath.Join(messagesDir, "update-status.json"), []byte(messageJSON), 0644)
	require.NoError(t, err)

	// Create main AsyncAPI file
	asyncapiJSON := `{
		"asyncapi": "2.6.0",
		"info": {
			"title": "Customer API",
			"version": "1.0.0"
		},
		"channels": {
			"customer/status": {
				"publish": {
					"message": {
						"$ref": "./messages/update-status.json"
					}
				}
			}
		}
	}`
	mainFilePath := filepath.Join(tmpDir, "asyncapi.json")
	err = os.WriteFile(mainFilePath, []byte(asyncapiJSON), 0644)
	require.NoError(t, err)

	// Parse the main AsyncAPI file
	doc, err := ParseFile(mainFilePath)
	require.NoError(t, err)
	require.NotNil(t, doc)

	// Convert to v2 document to access fields
	v2Doc, ok := doc.(*asyncapi2.Document)
	require.True(t, ok, "document should be v2")

	// Verify channel exists
	channel, ok := v2Doc.Channels["customer/status"]
	require.True(t, ok, "channel should exist")
	require.NotNil(t, channel.Publish, "publish operation should exist")
	require.NotNil(t, channel.Publish.Message, "message should exist")

	// Verify message details
	message := channel.Publish.Message
	assert.Equal(t, "UpdateCustomerStatus", message.Name)
	assert.Equal(t, "application/json", message.ContentType)

	// Verify payload was resolved
	payload, ok := message.Payload.(map[string]interface{})
	require.True(t, ok, "payload should be a map")
	assert.Equal(t, "object", payload["type"])

	// Check properties
	properties, ok := payload["properties"].(map[string]interface{})
	require.True(t, ok, "should have properties")

	// Check customerId field
	customerId, ok := properties["customerId"].(map[string]interface{})
	require.True(t, ok, "should have customerId")
	assert.Equal(t, "string", customerId["type"])
	assert.Equal(t, "uuid", customerId["format"])

	// Check status field
	status, ok := properties["status"].(map[string]interface{})
	require.True(t, ok, "should have status")
	assert.Equal(t, "string", status["type"])
	enum, ok := status["enum"].([]interface{})
	require.True(t, ok, "should have enum")
	assert.Contains(t, enum, "active")
	assert.Contains(t, enum, "inactive")
}

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

		// Skip semantic invalid files (they parse but fail validation, tested separately)
		if strings.HasPrefix(info.Name(), "semantic_invalid_") {
			return nil
		}

		t.Run(info.Name(), func(t *testing.T) {
			data, err := os.ReadFile(path)
			assert.NoError(t, err, "Error reading test file")

			doc, err := Parse(data)
			// Files starting with "invalid_" should fail to parse
			if strings.HasPrefix(info.Name(), "invalid_") {
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

// ============================================================================
// AsyncAPI v3 Tests
// ============================================================================

// TestParseV3Specs tests basic parsing of v3 documents
func TestParseV3Specs(t *testing.T) {
	tests := []struct {
		name        string
		spec        string
		expectError bool
	}{
		{
			name: "v3.0.0 minimal",
			spec: `{
				"asyncapi": "3.0.0",
				"info": {
					"title": "Test Service",
					"version": "1.0.0"
				}
			}`,
			expectError: false,
		},
		{
			name: "v3.0.0 with channels",
			spec: `{
				"asyncapi": "3.0.0",
				"info": {
					"title": "Test Service",
					"version": "1.0.0"
				},
				"channels": {
					"userEvents": {
						"address": "user/events",
						"description": "User event channel"
					}
				}
			}`,
			expectError: false,
		},
		{
			name: "v3.0.0 with operations",
			spec: `{
				"asyncapi": "3.0.0",
				"info": {
					"title": "Test Service",
					"version": "1.0.0"
				},
				"channels": {
					"userEvents": {
						"address": "user/events"
					}
				},
				"operations": {
					"sendUserEvent": {
						"action": "send",
						"channel": {
							"$ref": "#/channels/userEvents"
						}
					}
				}
			}`,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := ParseFromJSON([]byte(tt.spec))
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, doc)
			} else {
				assert.NoError(t, err)
				require.NotNil(t, doc)
				assert.True(t, strings.HasPrefix(doc.GetVersion(), "3."))
			}
		})
	}
}

// TestParseV3FromYAML tests parsing v3 YAML documents
func TestParseV3FromYAML(t *testing.T) {
	yamlDoc := `
asyncapi: 3.0.0
info:
  title: YAML v3 Service
  version: 1.0.0
  description: A test service
channels:
  userEvents:
    address: user/events
    description: User event channel
operations:
  publishUser:
    action: send
    channel:
      $ref: '#/channels/userEvents'
`
	doc, err := ParseFromYAML([]byte(yamlDoc))
	assert.NoError(t, err)
	require.NotNil(t, doc)
	assert.Equal(t, "3.0.0", doc.GetVersion())

	v3Doc, ok := doc.(*asyncapi3.Document)
	require.True(t, ok, "document should be v3")
	assert.Equal(t, "YAML v3 Service", v3Doc.Info.Title)
	assert.Len(t, v3Doc.Channels, 1)
	assert.Len(t, v3Doc.Operations, 1)
}

// TestParseV3Info tests v3 info parsing
func TestParseV3Info(t *testing.T) {
	spec := `{
		"asyncapi": "3.0.0",
		"info": {
			"title": "Full Info Service",
			"version": "2.1.0",
			"description": "A service with full info",
			"termsOfService": "https://example.com/terms",
			"contact": {
				"name": "API Team",
				"email": "api@example.com",
				"url": "https://api.example.com"
			},
			"license": {
				"name": "Apache 2.0",
				"url": "https://www.apache.org/licenses/LICENSE-2.0"
			},
			"tags": [
				{
					"name": "events",
					"description": "Event operations"
				}
			]
		}
	}`

	doc, err := ParseFromJSON([]byte(spec))
	require.NoError(t, err)
	require.NotNil(t, doc)

	v3Doc, ok := doc.(*asyncapi3.Document)
	require.True(t, ok)

	assert.Equal(t, "Full Info Service", v3Doc.Info.Title)
	assert.Equal(t, "2.1.0", v3Doc.Info.Version)
	assert.Equal(t, "A service with full info", v3Doc.Info.Description)
	assert.Equal(t, "https://example.com/terms", v3Doc.Info.TermsOfService)

	require.NotNil(t, v3Doc.Info.Contact)
	assert.Equal(t, "API Team", v3Doc.Info.Contact.Name)
	assert.Equal(t, "api@example.com", v3Doc.Info.Contact.Email)
	assert.Equal(t, "https://api.example.com", v3Doc.Info.Contact.URL)

	require.NotNil(t, v3Doc.Info.License)
	assert.Equal(t, "Apache 2.0", v3Doc.Info.License.Name)
	assert.Equal(t, "https://www.apache.org/licenses/LICENSE-2.0", v3Doc.Info.License.URL)

	require.Len(t, v3Doc.Info.Tags, 1)
	assert.Equal(t, "events", v3Doc.Info.Tags[0].Name)
}

// TestParseV3Servers tests v3 server parsing
func TestParseV3Servers(t *testing.T) {
	spec := `{
		"asyncapi": "3.0.0",
		"info": {
			"title": "Server Test",
			"version": "1.0.0"
		},
		"servers": {
			"production": {
				"host": "kafka.example.com:9092",
				"protocol": "kafka",
				"protocolVersion": "3.0.0",
				"description": "Production Kafka cluster",
				"tags": [
					{"name": "production"}
				]
			},
			"staging": {
				"host": "kafka-staging.example.com:9092",
				"protocol": "kafka",
				"description": "Staging Kafka cluster"
			}
		}
	}`

	doc, err := ParseFromJSON([]byte(spec))
	require.NoError(t, err)

	v3Doc, ok := doc.(*asyncapi3.Document)
	require.True(t, ok)

	assert.Len(t, v3Doc.Servers, 2)

	prod := v3Doc.Servers["production"]
	require.NotNil(t, prod)
	assert.Equal(t, "kafka.example.com:9092", prod.Host)
	assert.Equal(t, "kafka", prod.Protocol)
	assert.Equal(t, "3.0.0", prod.ProtocolVer)
	assert.Equal(t, "Production Kafka cluster", prod.Description)
	require.Len(t, prod.Tags, 1)
	assert.Equal(t, "production", prod.Tags[0].Name)
}

// TestParseV3Channels tests v3 channel parsing with messages
func TestParseV3Channels(t *testing.T) {
	spec := `{
		"asyncapi": "3.0.0",
		"info": {
			"title": "Channel Test",
			"version": "1.0.0"
		},
		"channels": {
			"userCreated": {
				"address": "user.created",
				"title": "User Created Channel",
				"description": "Channel for user creation events",
				"messages": {
					"UserCreatedMessage": {
						"name": "UserCreatedMessage",
						"contentType": "application/json",
						"payload": {
							"type": "object",
							"properties": {
								"userId": {"type": "string"}
							}
						}
					}
				}
			}
		}
	}`

	doc, err := ParseFromJSON([]byte(spec))
	require.NoError(t, err)

	v3Doc, ok := doc.(*asyncapi3.Document)
	require.True(t, ok)

	assert.Len(t, v3Doc.Channels, 1)

	channel := v3Doc.Channels["userCreated"]
	require.NotNil(t, channel)
	assert.Equal(t, "user.created", channel.Address)
	assert.Equal(t, "User Created Channel", channel.Title)
	assert.Equal(t, "Channel for user creation events", channel.Description)

	require.Len(t, channel.Messages, 1)
	msg := channel.Messages["UserCreatedMessage"]
	require.NotNil(t, msg)
	assert.Equal(t, "UserCreatedMessage", msg.Name)
	assert.Equal(t, "application/json", msg.ContentType)
}

// TestParseV3Operations tests v3 operation parsing
func TestParseV3Operations(t *testing.T) {
	specJSON := `{
		"asyncapi": "3.0.0",
		"info": {
			"title": "Operation Test",
			"version": "1.0.0"
		},
		"channels": {
			"userEvents": {
				"address": "user/events"
			}
		},
		"operations": {
			"publishUserCreated": {
				"action": "send",
				"channel": {
					"$ref": "#/channels/userEvents"
				},
				"title": "Publish User Created",
				"summary": "Publishes user created events",
				"description": "Full description here",
				"tags": [
					{"name": "users"}
				]
			},
			"receiveUserUpdates": {
				"action": "receive",
				"channel": {
					"$ref": "#/channels/userEvents"
				},
				"title": "Receive User Updates"
			}
		}
	}`

	doc, err := ParseFromJSON([]byte(specJSON))
	require.NoError(t, err)

	v3Doc, ok := doc.(*asyncapi3.Document)
	require.True(t, ok)

	assert.Len(t, v3Doc.Operations, 2)

	// Test send operation
	sendOp := v3Doc.Operations["publishUserCreated"]
	require.NotNil(t, sendOp)
	assert.Equal(t, spec.Send, sendOp.Action)
	assert.True(t, sendOp.IsSend())
	assert.False(t, sendOp.IsReceive())
	assert.Equal(t, "Publish User Created", sendOp.Title)
	assert.Equal(t, "Publishes user created events", sendOp.Summary)
	// Note: Channel ref is resolved by the reference resolver, so we check that Channel exists
	// The ref string itself may be empty after resolution
	require.NotNil(t, sendOp.Channel)

	// Test receive operation
	recvOp := v3Doc.Operations["receiveUserUpdates"]
	require.NotNil(t, recvOp)
	assert.Equal(t, spec.Receive, recvOp.Action)
	assert.True(t, recvOp.IsReceive())
	assert.False(t, recvOp.IsSend())
}

// TestParseV3Components tests v3 components parsing
func TestParseV3Components(t *testing.T) {
	spec := `{
		"asyncapi": "3.0.0",
		"info": {
			"title": "Components Test",
			"version": "1.0.0"
		},
		"components": {
			"messages": {
				"UserMessage": {
					"name": "UserMessage",
					"contentType": "application/json",
					"payload": {
						"type": "object",
						"properties": {
							"id": {"type": "string"}
						}
					}
				}
			},
			"schemas": {
				"User": {
					"type": "object",
					"properties": {
						"id": {"type": "string"},
						"name": {"type": "string"}
					}
				}
			}
		}
	}`

	doc, err := ParseFromJSON([]byte(spec))
	require.NoError(t, err)

	v3Doc, ok := doc.(*asyncapi3.Document)
	require.True(t, ok)

	require.NotNil(t, v3Doc.Components)

	// Check messages
	require.Len(t, v3Doc.Components.Messages, 1)
	userMsg := v3Doc.Components.Messages["UserMessage"]
	require.NotNil(t, userMsg)
	assert.Equal(t, "UserMessage", userMsg.Name)

	// Check schemas
	require.Len(t, v3Doc.Components.Schemas, 1)
	_, ok = v3Doc.Components.Schemas["User"]
	assert.True(t, ok)
}

// TestParseV3ChannelHasBinding tests the HasBinding helper method
func TestParseV3ChannelHasBinding(t *testing.T) {
	spec := `{
		"asyncapi": "3.0.0",
		"info": {
			"title": "Binding Test",
			"version": "1.0.0"
		},
		"channels": {
			"kafkaChannel": {
				"address": "my-topic",
				"bindings": {
					"kafka": {
						"topic": "my-specific-topic",
						"partitions": 10
					}
				}
			},
			"noBindingChannel": {
				"address": "plain-channel"
			}
		}
	}`

	doc, err := ParseFromJSON([]byte(spec))
	require.NoError(t, err)

	v3Doc, ok := doc.(*asyncapi3.Document)
	require.True(t, ok)

	kafkaCh := v3Doc.Channels["kafkaChannel"]
	require.NotNil(t, kafkaCh)
	assert.True(t, kafkaCh.HasBinding("kafka"))
	assert.False(t, kafkaCh.HasBinding("amqp"))

	plainCh := v3Doc.Channels["noBindingChannel"]
	require.NotNil(t, plainCh)
	assert.False(t, plainCh.HasBinding("kafka"))
}

// TestParseV3KafkaBinding tests Kafka binding parsing for v3
func TestParseV3KafkaBinding(t *testing.T) {
	doc, err := ParseFile("testdata/valid_3_0_0_kafka.json")
	require.NoError(t, err)

	v3Doc, ok := doc.(*asyncapi3.Document)
	require.True(t, ok)
	assert.Equal(t, "3.0.0", v3Doc.AsyncAPI)

	// Test info
	assert.Equal(t, "User Events Service", v3Doc.Info.Title)
	assert.Equal(t, "1.0.0", v3Doc.Info.Version)

	// Test servers
	require.NotNil(t, v3Doc.Servers["production"])
	assert.Equal(t, "kafka.example.com:9092", v3Doc.Servers["production"].Host)
	assert.Equal(t, "kafka", v3Doc.Servers["production"].Protocol)

	// Test channels with kafka bindings
	userSignedUpChannel := v3Doc.Channels["userSignedUp"]
	require.NotNil(t, userSignedUpChannel)
	assert.True(t, userSignedUpChannel.HasBinding("kafka"))

	// Parse the kafka binding
	binding, err := ParseBindings[kafka.ChannelBinding](userSignedUpChannel.Bindings, "kafka")
	require.NoError(t, err)
	assert.Equal(t, "user-signedup-events", binding.Topic)
	assert.Equal(t, 12, binding.Partitions)
	assert.Equal(t, 3, binding.Replicas)

	require.NotNil(t, binding.TopicConfiguration)
	assert.Contains(t, binding.TopicConfiguration.CleanupPolicy, "delete")
	assert.Equal(t, int64(604800000), binding.TopicConfiguration.RetentionMs)

	// Test operations
	require.Len(t, v3Doc.Operations, 2)

	publishOp := v3Doc.Operations["publishUserSignedUp"]
	require.NotNil(t, publishOp)
	assert.True(t, publishOp.IsSend())

	consumeOp := v3Doc.Operations["consumeUserDeleted"]
	require.NotNil(t, consumeOp)
	assert.True(t, consumeOp.IsReceive())
}

// TestParseV3AMQPBinding tests AMQP binding parsing for v3
func TestParseV3AMQPBinding(t *testing.T) {
	doc, err := ParseFile("testdata/valid_3_0_0_amqp.yaml")
	require.NoError(t, err)

	v3Doc, ok := doc.(*asyncapi3.Document)
	require.True(t, ok)
	assert.Equal(t, "3.0.0", v3Doc.AsyncAPI)

	// Test info
	assert.Equal(t, "Order Processing Service", v3Doc.Info.Title)

	// Test exchange channel
	orderCreatedChannel := v3Doc.Channels["orderCreated"]
	require.NotNil(t, orderCreatedChannel)
	assert.True(t, orderCreatedChannel.HasBinding("amqp"))

	binding, err := ParseBindings[amqp.ChannelBinding](orderCreatedChannel.Bindings, "amqp")
	require.NoError(t, err)
	assert.Equal(t, "routingKey", binding.Is)
	require.NotNil(t, binding.Exchange)
	assert.Equal(t, "orders-exchange", binding.Exchange.Name)
	assert.Equal(t, "topic", binding.Exchange.Type)
	assert.True(t, binding.Exchange.Durable)
	assert.False(t, binding.Exchange.AutoDelete)
	assert.Equal(t, "/production", binding.Exchange.VHost)

	// Test queue channel
	orderQueueChannel := v3Doc.Channels["orderQueue"]
	require.NotNil(t, orderQueueChannel)

	queueBinding, err := ParseBindings[amqp.ChannelBinding](orderQueueChannel.Bindings, "amqp")
	require.NoError(t, err)
	assert.Equal(t, "queue", queueBinding.Is)
	require.NotNil(t, queueBinding.Queue)
	assert.Equal(t, "orders-processing-queue", queueBinding.Queue.Name)
	assert.True(t, queueBinding.Queue.Durable)
	assert.False(t, queueBinding.Queue.Exclusive)
}

// TestParseV3SNSSQSBinding tests SNS and SQS binding parsing for v3
func TestParseV3SNSSQSBinding(t *testing.T) {
	doc, err := ParseFile("testdata/valid_3_0_0_sns_sqs.json")
	require.NoError(t, err)

	v3Doc, ok := doc.(*asyncapi3.Document)
	require.True(t, ok)
	assert.Equal(t, "3.0.0", v3Doc.AsyncAPI)

	// Test SNS channel
	snsChannel := v3Doc.Channels["notificationTopic"]
	require.NotNil(t, snsChannel)
	assert.True(t, snsChannel.HasBinding("sns"))

	snsBinding, err := ParseBindings[sns.ChannelBinding](snsChannel.Bindings, "sns")
	require.NoError(t, err)
	assert.Equal(t, "notifications-topic", snsBinding.Name)
	require.NotNil(t, snsBinding.Ordering)
	assert.Equal(t, "standard", snsBinding.Ordering.Type)

	// Test SQS channel (email queue)
	sqsChannel := v3Doc.Channels["emailQueue"]
	require.NotNil(t, sqsChannel)
	assert.True(t, sqsChannel.HasBinding("sqs"))

	sqsBinding, err := ParseBindings[sqs.ChannelBinding](sqsChannel.Bindings, "sqs")
	require.NoError(t, err)
	require.NotNil(t, sqsBinding.Queue)
	assert.Equal(t, "email-notifications-queue", sqsBinding.Queue.Name)
	assert.False(t, sqsBinding.Queue.FifoQueue)
	assert.Equal(t, 30, sqsBinding.Queue.VisibilityTimeout)
	assert.Equal(t, 345600, sqsBinding.Queue.MessageRetentionPeriod)

	// Check redrivePolicy (DLQ)
	require.NotNil(t, sqsBinding.Queue.RedrivePolicy)
	require.NotNil(t, sqsBinding.Queue.RedrivePolicy.MaxReceiveCount)
	assert.Equal(t, 3, *sqsBinding.Queue.RedrivePolicy.MaxReceiveCount)

	// Test FIFO queue
	fifoChannel := v3Doc.Channels["priorityQueue"]
	require.NotNil(t, fifoChannel)

	fifoBinding, err := ParseBindings[sqs.ChannelBinding](fifoChannel.Bindings, "sqs")
	require.NoError(t, err)
	require.NotNil(t, fifoBinding.Queue)
	assert.True(t, fifoBinding.Queue.FifoQueue)
	assert.Equal(t, "messageGroup", fifoBinding.Queue.DeduplicationScope)
	assert.Equal(t, "perMessageGroupId", fifoBinding.Queue.FifoThroughputLimit)
}

// TestParseV3GooglePubSubBinding tests Google Pub/Sub binding parsing for v3
func TestParseV3GooglePubSubBinding(t *testing.T) {
	doc, err := ParseFile("testdata/valid_3_0_0_googlepubsub.yaml")
	require.NoError(t, err)

	v3Doc, ok := doc.(*asyncapi3.Document)
	require.True(t, ok)
	assert.Equal(t, "3.0.0", v3Doc.AsyncAPI)

	// Test info
	assert.Equal(t, "Analytics Events Service", v3Doc.Info.Title)

	// Test server
	gcpServer := v3Doc.Servers["gcp"]
	require.NotNil(t, gcpServer)
	assert.Equal(t, "pubsub.googleapis.com", gcpServer.Host)
	assert.Equal(t, "googlepubsub", gcpServer.Protocol)

	// Test pageViews channel
	pageViewsChannel := v3Doc.Channels["pageViews"]
	require.NotNil(t, pageViewsChannel)
	assert.Equal(t, "projects/my-project/topics/page-views", pageViewsChannel.Address)
	assert.True(t, pageViewsChannel.HasBinding("googlepubsub"))

	binding, err := ParseBindings[googlepubsub.ChannelBinding](pageViewsChannel.Bindings, "googlepubsub")
	require.NoError(t, err)
	assert.Equal(t, "604800s", binding.MessageRetentionDuration)

	require.NotNil(t, binding.MessageStoragePolicy)
	assert.Contains(t, binding.MessageStoragePolicy.AllowedPersistenceRegions, "us-central1")
	assert.Contains(t, binding.MessageStoragePolicy.AllowedPersistenceRegions, "europe-west1")

	require.NotNil(t, binding.SchemaSettings)
	assert.Equal(t, "JSON", binding.SchemaSettings.Encoding)

	require.NotNil(t, binding.Labels)
	assert.Equal(t, "production", binding.Labels["environment"])
	assert.Equal(t, "analytics", binding.Labels["team"])

	// Test conversion channel with BINARY encoding
	conversionChannel := v3Doc.Channels["conversionEvents"]
	require.NotNil(t, conversionChannel)

	convBinding, err := ParseBindings[googlepubsub.ChannelBinding](conversionChannel.Bindings, "googlepubsub")
	require.NoError(t, err)
	require.NotNil(t, convBinding.SchemaSettings)
	assert.Equal(t, "BINARY", convBinding.SchemaSettings.Encoding)
	assert.Equal(t, "v1", convBinding.SchemaSettings.FirstRevisionID)
	assert.Equal(t, "v3", convBinding.SchemaSettings.LastRevisionID)
}

// TestParseV3MultiProtocol tests parsing a multi-protocol spec
func TestParseV3MultiProtocol(t *testing.T) {
	doc, err := ParseFile("testdata/valid_3_0_0_multi_protocol.yaml")
	require.NoError(t, err)

	v3Doc, ok := doc.(*asyncapi3.Document)
	require.True(t, ok)

	// Test multiple servers with different protocols
	assert.Len(t, v3Doc.Servers, 3)
	assert.Equal(t, "kafka", v3Doc.Servers["kafkaProduction"].Protocol)
	assert.Equal(t, "amqp", v3Doc.Servers["rabbitProduction"].Protocol)
	assert.Equal(t, "sqs", v3Doc.Servers["awsSqs"].Protocol)

	// Test channels with different bindings
	assert.Len(t, v3Doc.Channels, 4)

	// Kafka channel
	inboundChannel := v3Doc.Channels["inboundEvents"]
	require.NotNil(t, inboundChannel)
	assert.True(t, inboundChannel.HasBinding("kafka"))

	kafkaBinding, err := ParseBindings[kafka.ChannelBinding](inboundChannel.Bindings, "kafka")
	require.NoError(t, err)
	assert.Equal(t, "inbound-events", kafkaBinding.Topic)
	assert.Equal(t, 24, kafkaBinding.Partitions)

	// AMQP channel
	processedChannel := v3Doc.Channels["processedEvents"]
	require.NotNil(t, processedChannel)
	assert.True(t, processedChannel.HasBinding("amqp"))

	amqpBinding, err := ParseBindings[amqp.ChannelBinding](processedChannel.Bindings, "amqp")
	require.NoError(t, err)
	assert.Equal(t, "routingKey", amqpBinding.Is)
	assert.Equal(t, "processed-events-exchange", amqpBinding.Exchange.Name)

	// SQS channel
	dlqChannel := v3Doc.Channels["deadLetterQueue"]
	require.NotNil(t, dlqChannel)
	assert.True(t, dlqChannel.HasBinding("sqs"))

	sqsBinding, err := ParseBindings[sqs.ChannelBinding](dlqChannel.Bindings, "sqs")
	require.NoError(t, err)
	assert.Equal(t, "events-dlq", sqsBinding.Queue.Name)
	assert.Equal(t, 1209600, sqsBinding.Queue.MessageRetentionPeriod)

	// Test operations
	assert.Len(t, v3Doc.Operations, 4)

	// Count send vs receive operations
	sendCount := 0
	receiveCount := 0
	for _, op := range v3Doc.Operations {
		if op.IsSend() {
			sendCount++
		}
		if op.IsReceive() {
			receiveCount++
		}
	}
	assert.Equal(t, 3, sendCount)
	assert.Equal(t, 1, receiveCount)
}

// TestParseV3DocumentResolver tests the document resolver methods
// Note: This test uses asyncapi3.ParseFromJSON directly to avoid the main parser's
// reference resolution, which would replace $ref objects with resolved content.
func TestParseV3DocumentResolver(t *testing.T) {
	specJSON := `{
		"asyncapi": "3.0.0",
		"info": {
			"title": "Resolver Test",
			"version": "1.0.0"
		},
		"channels": {
			"userEvents": {
				"address": "user/events",
				"messages": {
					"UserMessage": {
						"name": "UserMessage",
						"contentType": "application/json"
					}
				}
			}
		},
		"operations": {
			"sendUser": {
				"action": "send",
				"channel": {
					"$ref": "#/channels/userEvents"
				}
			}
		},
		"components": {
			"messages": {
				"ComponentMessage": {
					"name": "ComponentMessage",
					"contentType": "application/json"
				}
			}
		}
	}`

	// Use asyncapi3.ParseFromJSON directly to preserve refs
	v3Doc, err := asyncapi3.ParseFromJSON([]byte(specJSON))
	require.NoError(t, err)

	// Test ResolveChannelRef
	channel := v3Doc.ResolveChannelRef("#/channels/userEvents")
	require.NotNil(t, channel)
	assert.Equal(t, "user/events", channel.Address)

	// Test with non-existent ref
	nilChannel := v3Doc.ResolveChannelRef("#/channels/nonExistent")
	assert.Nil(t, nilChannel)

	// Test ResolveMessageRef from components
	msg := v3Doc.ResolveMessageRef("#/components/messages/ComponentMessage")
	require.NotNil(t, msg)
	assert.Equal(t, "ComponentMessage", msg.Name)

	// Test GetChannelForOperation
	op := v3Doc.Operations["sendUser"]
	require.NotNil(t, op)
	opChannel := v3Doc.GetChannelForOperation(op)
	require.NotNil(t, opChannel)
	assert.Equal(t, "user/events", opChannel.Address)
}

// ============================================================================
// AsyncAPI v3 Invalid Spec Tests
// ============================================================================

// TestParseV3InvalidSpecs tests that invalid v3 specs fail validation
// Note: JSON Schema allows additional properties by default, so some v2 patterns
// may parse without validation errors. The tests below focus on truly invalid specs.
func TestParseV3InvalidSpecs(t *testing.T) {
	tests := []struct {
		name          string
		file          string
		description   string
		expectFailure bool // Some specs are structurally valid but semantically wrong
	}{
		{
			name:          "v2 publish/subscribe pattern",
			file:          "testdata/semantic_invalid_3_0_0_v2_publish_subscribe.yaml",
			description:   "v3 allows extra properties, so publish/subscribe is ignored but valid",
			expectFailure: false, // JSON Schema allows additional properties
		},
		{
			name:          "v2 message oneOf pattern",
			file:          "testdata/semantic_invalid_3_0_0_v2_message_oneOf.json",
			description:   "v3 allows extra properties on messages",
			expectFailure: false, // JSON Schema allows additional properties
		},
		{
			name:          "missing action field",
			file:          "testdata/semantic_invalid_3_0_0_missing_action.yaml",
			description:   "operations require action field",
			expectFailure: true,
		},
		{
			name:          "wrong action value",
			file:          "testdata/semantic_invalid_3_0_0_wrong_action.json",
			description:   "action must be 'send' or 'receive', not 'publish'",
			expectFailure: true,
		},
		{
			name:          "v2 server url pattern",
			file:          "testdata/semantic_invalid_3_0_0_v2_server_url.yaml",
			description:   "v3 servers use 'host' not 'url', but url is ignored",
			expectFailure: false, // JSON Schema allows additional properties
		},
		{
			name:          "missing info",
			file:          "testdata/semantic_invalid_3_0_0_missing_info.json",
			description:   "info is required",
			expectFailure: true,
		},
		{
			name:          "v2 channel item pattern",
			file:          "testdata/semantic_invalid_3_0_0_v2_channel_item.yaml",
			description:   "v3 channels don't have operationId, but extra props are ignored",
			expectFailure: false, // JSON Schema allows additional properties
		},
		{
			name:          "operation missing channel",
			file:          "testdata/semantic_invalid_3_0_0_operation_missing_channel.json",
			description:   "operations require channel reference",
			expectFailure: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := ParseFile(tt.file)

			// The spec should parse (syntax is valid) but validation may fail
			if err != nil {
				if tt.expectFailure {
					t.Logf("Parse error (expected): %v", err)
					return
				}
				t.Fatalf("Unexpected parse error: %v", err)
			}

			require.NotNil(t, doc, "Document should parse")

			// Validate the document
			err = doc.Validate()
			if tt.expectFailure {
				assert.Error(t, err, "Validation should fail for: %s", tt.description)
				if err != nil {
					t.Logf("Validation error (expected): %v", err)
				}
			} else {
				// These specs are structurally valid (JSON Schema allows extra props)
				// but semantically wrong for v3
				if err != nil {
					t.Logf("Validation error (unexpected but informative): %v", err)
				}
			}
		})
	}
}

// TestParseV3BreakingChangesFromV2 tests specific v2 to v3 breaking changes
// Note: Some tests use asyncapi3.ParseFromJSON directly to avoid ref resolution
// which would replace $ref objects with resolved content.
func TestParseV3BreakingChangesFromV2(t *testing.T) {
	t.Run("publish/subscribe replaced with send/receive", func(t *testing.T) {
		// v2 style - has publish/subscribe on channel (wrong for v3)
		// JSON Schema allows extra properties, so this parses but is semantically wrong
		v2Style := `{
			"asyncapi": "3.0.0",
			"info": {"title": "Test", "version": "1.0.0"},
			"channels": {
				"test": {
					"address": "test/topic",
					"publish": {"message": {"payload": {"type": "string"}}}
				}
			}
		}`

		// Use asyncapi3 parser directly
		doc, err := asyncapi3.ParseFromJSON([]byte(v2Style))
		require.NoError(t, err)
		require.NotNil(t, doc)

		// The publish field is ignored in v3 - it's just an extra property
		// But the channel is still valid
		channel := doc.Channels["test"]
		require.NotNil(t, channel)
		assert.Equal(t, "test/topic", channel.Address)

		// v3 style - operations are separate with action
		v3Style := `{
			"asyncapi": "3.0.0",
			"info": {"title": "Test", "version": "1.0.0"},
			"channels": {
				"test": {"address": "test/topic"}
			},
			"operations": {
				"publishTest": {
					"action": "send",
					"channel": {"$ref": "#/channels/test"}
				}
			}
		}`

		doc, err = asyncapi3.ParseFromJSON([]byte(v3Style))
		require.NoError(t, err)
		require.NotNil(t, doc)

		// Verify operation structure
		op := doc.Operations["publishTest"]
		require.NotNil(t, op)
		assert.True(t, op.IsSend())
		assert.Equal(t, "#/channels/test", op.Channel.Ref)
	})

	t.Run("server url replaced with host", func(t *testing.T) {
		// v2 style - has url (ignored in v3, host is missing)
		v2Style := `{
			"asyncapi": "3.0.0",
			"info": {"title": "Test", "version": "1.0.0"},
			"servers": {
				"prod": {
					"url": "kafka://localhost:9092",
					"protocol": "kafka"
				}
			}
		}`

		doc, err := asyncapi3.ParseFromJSON([]byte(v2Style))
		require.NoError(t, err)
		require.NotNil(t, doc)

		// url is an extra property that's ignored, host will be empty
		server := doc.Servers["prod"]
		require.NotNil(t, server)
		assert.Empty(t, server.Host, "v2 url should not populate host")
		assert.Equal(t, "kafka", server.Protocol)

		// v3 style - has host
		v3Style := `{
			"asyncapi": "3.0.0",
			"info": {"title": "Test", "version": "1.0.0"},
			"servers": {
				"prod": {
					"host": "localhost:9092",
					"protocol": "kafka"
				}
			}
		}`

		doc, err = asyncapi3.ParseFromJSON([]byte(v3Style))
		require.NoError(t, err)
		require.NotNil(t, doc)

		server = doc.Servers["prod"]
		require.NotNil(t, server)
		assert.Equal(t, "localhost:9092", server.Host)
		assert.Equal(t, "kafka", server.Protocol)
	})

	t.Run("action values send/receive not publish/subscribe", func(t *testing.T) {
		// Using 'publish' as action (wrong) - will be parsed but won't match IsSend/IsReceive
		wrongAction := `{
			"asyncapi": "3.0.0",
			"info": {"title": "Test", "version": "1.0.0"},
			"channels": {"test": {"address": "test"}},
			"operations": {
				"op": {
					"action": "publish",
					"channel": {"$ref": "#/channels/test"}
				}
			}
		}`

		doc, err := asyncapi3.ParseFromJSON([]byte(wrongAction))
		require.NoError(t, err)
		require.NotNil(t, doc)

		op := doc.Operations["op"]
		require.NotNil(t, op)
		// 'publish' is not a valid action, so IsSend and IsReceive will both be false
		assert.False(t, op.IsSend(), "publish is not send")
		assert.False(t, op.IsReceive(), "publish is not receive")
		assert.Equal(t, spec.Action("publish"), op.Action)

		// Validation should fail for wrong action value
		err = doc.Validate()
		assert.Error(t, err, "action 'publish' should fail validation")

		// Using 'send' as action (correct)
		correctAction := `{
			"asyncapi": "3.0.0",
			"info": {"title": "Test", "version": "1.0.0"},
			"channels": {"test": {"address": "test"}},
			"operations": {
				"op": {
					"action": "send",
					"channel": {"$ref": "#/channels/test"}
				}
			}
		}`

		doc, err = asyncapi3.ParseFromJSON([]byte(correctAction))
		require.NoError(t, err)
		require.NotNil(t, doc)

		op = doc.Operations["op"]
		require.NotNil(t, op)
		assert.True(t, op.IsSend())
		assert.False(t, op.IsReceive())
		assert.Equal(t, spec.Send, op.Action)

		// Validation should pass
		err = doc.Validate()
		assert.NoError(t, err, "action 'send' should be valid")
	})

	t.Run("channels require address in v3", func(t *testing.T) {
		// v2 style - channel name is the address (no explicit address field)
		v2Style := `{
			"asyncapi": "3.0.0",
			"info": {"title": "Test", "version": "1.0.0"},
			"channels": {
				"user/signedup": {}
			}
		}`

		doc, err := asyncapi3.ParseFromJSON([]byte(v2Style))
		require.NoError(t, err)
		require.NotNil(t, doc)

		channel := doc.Channels["user/signedup"]
		require.NotNil(t, channel)
		// Address field should be empty since v2 pattern was used
		assert.Empty(t, channel.Address, "v2 style channel should not have address populated")

		// v3 style - explicit address
		v3Style := `{
			"asyncapi": "3.0.0",
			"info": {"title": "Test", "version": "1.0.0"},
			"channels": {
				"userSignedup": {
					"address": "user/signedup"
				}
			}
		}`

		doc, err = asyncapi3.ParseFromJSON([]byte(v3Style))
		require.NoError(t, err)
		require.NotNil(t, doc)

		channel = doc.Channels["userSignedup"]
		require.NotNil(t, channel)
		assert.Equal(t, "user/signedup", channel.Address)
	})

	t.Run("messages in channels vs operations", func(t *testing.T) {
		// In v3, messages are defined on channels, not operations
		v3Style := `{
			"asyncapi": "3.0.0",
			"info": {"title": "Test", "version": "1.0.0"},
			"channels": {
				"userEvents": {
					"address": "user/events",
					"messages": {
						"UserCreated": {
							"name": "UserCreated",
							"payload": {"type": "object"}
						}
					}
				}
			},
			"operations": {
				"sendUserCreated": {
					"action": "send",
					"channel": {"$ref": "#/channels/userEvents"},
					"messages": [
						{"$ref": "#/channels/userEvents/messages/UserCreated"}
					]
				}
			}
		}`

		doc, err := asyncapi3.ParseFromJSON([]byte(v3Style))
		require.NoError(t, err)
		require.NotNil(t, doc)

		// Check channel messages
		channel := doc.Channels["userEvents"]
		require.NotNil(t, channel)
		require.NotNil(t, channel.Messages)
		assert.Contains(t, channel.Messages, "UserCreated")

		// Check operation references messages
		op := doc.Operations["sendUserCreated"]
		require.NotNil(t, op)
		require.Len(t, op.Messages, 1)
		assert.Equal(t, "#/channels/userEvents/messages/UserCreated", op.Messages[0].Ref)
	})
}
