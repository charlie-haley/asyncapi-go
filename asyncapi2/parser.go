package asyncapi2

import (
	"encoding/json"
	"fmt"

	"github.com/charlie-haley/asyncapi-go/spec"
	"sigs.k8s.io/yaml"
)

// ParseFromJSON parses an AsyncAPI v2 document from JSON
func ParseFromJSON(data []byte) (spec.Document, error) {
	var doc Document
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if err := doc.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	return &doc, nil
}

// ParseFromYAML parses an AsyncAPI v2 document from YAML
func ParseFromYAML(data []byte) (spec.Document, error) {
	jsonData, err := yaml.YAMLToJSON(data)
	if err != nil {
		return nil, fmt.Errorf("failed to convert YAML to JSON: %w", err)
	}
	return ParseFromJSON(jsonData)
}
