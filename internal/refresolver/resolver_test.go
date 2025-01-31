package refresolver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Schema represents a JSON Schema structure
type Schema struct {
	Type       string             `json:"type,omitempty"`
	Format     string             `json:"format,omitempty"`
	Properties map[string]*Schema `json:"properties,omitempty"`
	Ref        string             `json:"$ref,omitempty"`
}

// Message represents an AsyncAPI message
type Message struct {
	Payload *Schema `json:"payload,omitempty"`
	Ref     string  `json:"$ref,omitempty"`
}

// Operation represents an AsyncAPI operation (publish/subscribe)
type Operation struct {
	Message *Message `json:"message,omitempty"`
}

// Channel represents an AsyncAPI channel
type Channel struct {
	Publish *Operation `json:"publish,omitempty"`
}

// Components represents the components section of AsyncAPI
type Components struct {
	Messages map[string]*Message `json:"messages,omitempty"`
	Schemas  map[string]*Schema  `json:"schemas,omitempty"`
}

// Document represents a simplified AsyncAPI document
type Document struct {
	Components *Components         `json:"components,omitempty"`
	Channels   map[string]*Channel `json:"channels,omitempty"`
	Message    *Message            `json:"message,omitempty"`
	Schema     *Schema             `json:"schema,omitempty"`
}

func TestResolveLocalRef(t *testing.T) {
	doc := &Document{
		Components: &Components{
			Messages: map[string]*Message{
				"TestMessage": {
					Payload: &Schema{
						Type: "string",
					},
				},
			},
		},
		Channels: map[string]*Channel{
			"test": {
				Publish: &Operation{
					Message: &Message{
						Ref: "#/components/messages/TestMessage",
					},
				},
			},
		},
	}

	// Convert to map for resolver
	data, err := json.Marshal(doc)
	require.NoError(t, err)
	var docMap map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &docMap))

	resolver := New("")
	resolver.Cache["#"] = docMap

	resolved, err := resolver.ResolveRefs(docMap)
	require.NoError(t, err)

	// Convert back to Document
	data, err = json.Marshal(resolved)
	require.NoError(t, err)
	var resolvedDoc Document
	require.NoError(t, json.Unmarshal(data, &resolvedDoc))

	message := resolvedDoc.Channels["test"].Publish.Message
	assert.Equal(t, "string", message.Payload.Type)
}

func TestResolveFileRef(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.json")

	testSchema := &Schema{
		Type:   "string",
		Format: "email",
	}

	data, err := json.Marshal(testSchema)
	require.NoError(t, err)
	err = os.WriteFile(testFile, data, 0644)
	require.NoError(t, err)

	doc := &Document{
		Schema: &Schema{
			Ref: testFile,
		},
	}

	// Convert to map for resolver
	data, err = json.Marshal(doc)
	require.NoError(t, err)
	var docMap map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &docMap))

	resolver := New(tmpDir)
	resolved, err := resolver.ResolveRefs(docMap)
	require.NoError(t, err)

	// Convert back to Document
	data, err = json.Marshal(resolved)
	require.NoError(t, err)
	var resolvedDoc Document
	require.NoError(t, json.Unmarshal(data, &resolvedDoc))

	assert.Equal(t, "string", resolvedDoc.Schema.Type)
	assert.Equal(t, "email", resolvedDoc.Schema.Format)
}

func TestResolveRemoteRef(t *testing.T) {
	testSchema := &Schema{
		Type:   "string",
		Format: "email",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(testSchema)
	}))
	defer server.Close()

	doc := &Document{
		Schema: &Schema{
			Ref: server.URL + "/test.json",
		},
	}

	// Convert to map for resolver
	data, err := json.Marshal(doc)
	require.NoError(t, err)
	var docMap map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &docMap))

	resolver := New("")
	resolved, err := resolver.ResolveRefs(docMap)
	require.NoError(t, err)

	// Convert back to Document
	data, err = json.Marshal(resolved)
	require.NoError(t, err)
	var resolvedDoc Document
	require.NoError(t, json.Unmarshal(data, &resolvedDoc))

	assert.Equal(t, "string", resolvedDoc.Schema.Type)
	assert.Equal(t, "email", resolvedDoc.Schema.Format)
}

func TestCircularRef(t *testing.T) {
	doc := &Document{
		Components: &Components{
			Schemas: map[string]*Schema{
				"A": {Ref: "#/components/schemas/B"},
				"B": {Ref: "#/components/schemas/A"},
			},
		},
	}

	// Convert to map for resolver
	data, err := json.Marshal(doc)
	require.NoError(t, err)
	var docMap map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &docMap))

	resolver := New("")
	resolver.Cache["#"] = docMap

	_, err = resolver.ResolveRefs(docMap)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "circular reference")
}

func TestNestedRefs(t *testing.T) {
	doc := &Document{
		Components: &Components{
			Schemas: map[string]*Schema{
				"Address": {
					Type: "object",
					Properties: map[string]*Schema{
						"street": {
							Type: "string",
						},
					},
				},
				"User": {
					Type: "object",
					Properties: map[string]*Schema{
						"name": {
							Type: "string",
						},
						"address": {
							Ref: "#/components/schemas/Address",
						},
					},
				},
			},
		},
		Schema: &Schema{
			Ref: "#/components/schemas/User",
		},
	}

	// Convert to map for resolver
	data, err := json.Marshal(doc)
	require.NoError(t, err)
	var docMap map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &docMap))

	resolver := New("")
	resolver.Cache["#"] = docMap

	resolved, err := resolver.ResolveRefs(docMap)
	require.NoError(t, err)

	// Convert back to Document
	data, err = json.Marshal(resolved)
	require.NoError(t, err)
	var resolvedDoc Document
	require.NoError(t, json.Unmarshal(data, &resolvedDoc))

	// After resolution, check the resolved Schema
	require.NotNil(t, resolvedDoc.Schema)
	assert.Equal(t, "object", resolvedDoc.Schema.Type)
	require.NotNil(t, resolvedDoc.Schema.Properties)

	// Check User properties
	require.Contains(t, resolvedDoc.Schema.Properties, "name")
	nameSchema := resolvedDoc.Schema.Properties["name"]
	assert.Equal(t, "string", nameSchema.Type)

	// Check Address properties
	require.Contains(t, resolvedDoc.Schema.Properties, "address")
	address := resolvedDoc.Schema.Properties["address"]
	require.NotNil(t, address)
	assert.Equal(t, "object", address.Type)
	require.NotNil(t, address.Properties)
	require.Contains(t, address.Properties, "street")
	assert.Equal(t, "string", address.Properties["street"].Type)
}
