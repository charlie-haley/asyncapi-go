package asyncapi

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charlie-haley/asyncapi-go/asyncapi2"
	"github.com/charlie-haley/asyncapi-go/internal/refresolver"
	"github.com/charlie-haley/asyncapi-go/spec"
	"sigs.k8s.io/yaml"
)

// BindingUnmarshaler represents a binding that can unmarshal itself
type BindingUnmarshaler interface {
	UnmarshalJSON([]byte) error
	UnmarshalYAML(func(interface{}) error) error
}

// ParseBindings processes bindings for a given channel/operation/message
func ParseBindings[T any](rawBindings map[string]interface{}, bindingType string) (*T, error) {
	var binding T
	if raw, ok := rawBindings[bindingType]; ok {
		data, err := json.Marshal(raw)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal %s binding: %w", bindingType, err)
		}

		if err := json.Unmarshal(data, &binding); err != nil {
			return nil, fmt.Errorf("failed to unmarshal %s binding: %w", bindingType, err)
		}
		return &binding, nil
	}

	return nil, fmt.Errorf("binding type %s not found", bindingType) // Error if binding doesn't exist
}

// ParseFromJSON parses an AsyncAPI document from JSON
func ParseFromJSON(data []byte) (spec.Document, error) {
	// First resolve all references
	var jsonDoc interface{}
	if err := json.Unmarshal(data, &jsonDoc); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Create resolver with current directory as base path
	resolver := refresolver.New(filepath.Dir("."))
	resolver.Cache["#"] = jsonDoc

	// Resolve all references in the document
	resolvedDoc, err := resolver.ResolveRefs(jsonDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve references: %w", err)
	}

	// Convert back to JSON to parse as AsyncAPI document
	resolvedData, err := json.Marshal(resolvedDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resolved document: %w", err)
	}

	// Unmarshal just enough to get the version
	var versionDoc struct {
		Version string `json:"asyncapi"`
	}
	if err := json.Unmarshal(resolvedData, &versionDoc); err != nil {
		return nil, fmt.Errorf("failed to parse document version: %w", err)
	}

	// Parse according to version
	switch {
	case strings.HasPrefix(versionDoc.Version, "2."):
		doc, err := asyncapi2.ParseFromJSON(resolvedData)
		if err != nil {
			return nil, err
		}
		return spec.Document(doc), nil
	default:
		return nil, fmt.Errorf("unsupported AsyncAPI version: %s", versionDoc.Version)
	}
}

// ParseFromYAML parses an AsyncAPI document from YAML
func ParseFromYAML(data []byte) (spec.Document, error) {
	jsonData, err := yaml.YAMLToJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to convert YAML to JSON: %w", err)
	}
	return ParseFromJSON(jsonData)
}

// Parse detects format and parses accordingly
func Parse(data []byte) (spec.Document, error) {
	if isYAML(data) {
		return ParseFromYAML(data)
	}
	return ParseFromJSON(data)
}

// isYAML determines if the input appears to be YAML
func isYAML(data []byte) bool {
	trimmed := strings.TrimSpace(string(data))
	return !strings.HasPrefix(trimmed, "{") && !strings.HasPrefix(trimmed, "[")
}
