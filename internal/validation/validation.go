package validation

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/asyncapi/spec-json-schemas/v6"
	"github.com/xeipuuv/gojsonschema"
)

// ErrUnsupportedVersion indicates that the AsyncAPI version is not supported
var ErrUnsupportedVersion = fmt.Errorf("unsupported AsyncAPI version")

// ValidateDocument validates an AsyncAPI document against its schema
func ValidateDocument(doc interface{}) error {
	// Get version from document
	version := getVersionFromDoc(doc)
	if version == "" {
		return fmt.Errorf("could not determine AsyncAPI version from document")
	}

	// Get schema from spec-json-schemas
	schema, err := spec_json_schemas.Get(version)
	if err != nil {
		return fmt.Errorf("failed to get schema for version %s: %w", version, err)
	}

	// Marshal document to JSON for validation
	docBytes, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("failed to marshal document for validation: %w", err)
	}

	// Setup schema validation
	schemaLoader := gojsonschema.NewBytesLoader(schema)
	docLoader := gojsonschema.NewBytesLoader(docBytes)

	// Validate
	result, err := gojsonschema.Validate(schemaLoader, docLoader)
	if err != nil {
		return fmt.Errorf("schema validation failed: %w", err)
	}

	if !result.Valid() {
		var errMsgs []string
		for _, err := range result.Errors() {
			errMsgs = append(errMsgs, fmt.Sprintf("- %s: %s", err.Field(), err.Description()))
		}
		return fmt.Errorf("validation errors:\n%s", strings.Join(errMsgs, "\n"))
	}

	return nil
}

// getVersionFromDoc extracts the AsyncAPI version from a document
func getVersionFromDoc(doc interface{}) string {
	// Try to get version through interface method first
	if d, ok := doc.(interface{ GetVersion() string }); ok {
		return d.GetVersion()
	}

	// Fallback to type assertion for map (useful during initial parsing)
	if m, ok := doc.(map[string]interface{}); ok {
		if v, ok := m["asyncapi"].(string); ok {
			return v
		}
	}

	return ""
}
