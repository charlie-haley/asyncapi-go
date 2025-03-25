package asyncapi

import (
	"encoding/json"
	"fmt"
	"os"
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

// ParseOptions contains options for parsing AsyncAPI documents
type ParseOptions struct {
	// FilePath is the path to the file being parsed. This is used to resolve relative refs.
	FilePath string
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

	return nil, fmt.Errorf("binding type %s not found", bindingType)
}

// ParseFromJSON parses an AsyncAPI document from JSON
func ParseFromJSON(data []byte, opts ...ParseOptions) (spec.Document, error) {
	var jsonDoc interface{}
	if err := json.Unmarshal(data, &jsonDoc); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	basePath := "."
	if len(opts) > 0 && opts[0].FilePath != "" {
		basePath = filepath.Dir(opts[0].FilePath)
	}

	resolver := refresolver.New(basePath)
	resolver.Cache["#"] = jsonDoc

	resolvedDoc, err := resolver.ResolveRefs(jsonDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve references: %w", err)
	}

	resolvedData, err := json.Marshal(resolvedDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resolved document: %w", err)
	}

	var versionDoc struct {
		Version string `json:"asyncapi"`
	}
	if err := json.Unmarshal(resolvedData, &versionDoc); err != nil {
		return nil, fmt.Errorf("failed to parse document version: %w", err)
	}

	switch {
	case strings.HasPrefix(versionDoc.Version, "2."):
		doc, err := asyncapi2.ParseFromJSON(resolvedData)
		if err != nil {
			return nil, err
		}
		return doc, nil
	default:
		return nil, fmt.Errorf("unsupported AsyncAPI version: %s", versionDoc.Version)
	}
}

// ParseFromYAML parses an AsyncAPI document from YAML
func ParseFromYAML(data []byte, opts ...ParseOptions) (spec.Document, error) {
	jsonData, err := yaml.YAMLToJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to convert YAML to JSON: %w", err)
	}
	return ParseFromJSON(jsonData, opts...)
}

// Parse detects format and parses accordingly
func Parse(data []byte, opts ...ParseOptions) (spec.Document, error) {
	if isYAML(data) {
		return ParseFromYAML(data, opts...)
	}
	return ParseFromJSON(data, opts...)
}

// isYAML determines if the input appears to be YAML
func isYAML(data []byte) bool {
	trimmed := strings.TrimSpace(string(data))
	return !strings.HasPrefix(trimmed, "{") && !strings.HasPrefix(trimmed, "[")
}

// ParseFile reads and parses an AsyncAPI file, automatically handling the filepath for reference resolution
func ParseFile(filePath string) (spec.Document, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return Parse(data, ParseOptions{FilePath: filePath})
}