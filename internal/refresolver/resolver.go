package refresolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"sigs.k8s.io/yaml"
)

type RefResolver struct {
	Cache       map[string]interface{}
	basePath    string
	currentFile string
}

func New(basePath string) *RefResolver {
	return &RefResolver{
		Cache:    make(map[string]interface{}),
		basePath: basePath,
	}
}

func (r *RefResolver) ResolveRefs(doc interface{}) (interface{}, error) {
	return r.resolveRefsRecursive(doc, make(map[string]bool), false)
}

func (r *RefResolver) resolveRefsRecursive(v interface{}, visited map[string]bool, isMessage bool) (interface{}, error) {
	switch val := v.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})

		// Check if this is a message with schema and headers
		if _, hasPayload := val["payload"]; hasPayload {
			if schemaFormat, ok := val["schemaFormat"].(string); ok && strings.Contains(schemaFormat, "application/schema+yaml") {
				// Deep resolve all refs in the message
				for k, v := range val {
					if k == "payload" || k == "headers" {
						if m, ok := v.(map[string]interface{}); ok {
							if ref, ok := m["$ref"]; ok {
								refStr, ok := ref.(string)
								if !ok {
									return nil, fmt.Errorf("$ref value must be a string, got %T", ref)
								}
								resolved, err := r.resolveRef(refStr)
								if err != nil {
									return nil, err
								}
								result[k] = resolved
								continue
							}
						}
					}
					resolved, err := r.resolveRefsRecursive(v, visited, false)
					if err != nil {
						return nil, err
					}
					result[k] = resolved
				}
				return result, nil
			}
		}

		// Handle other object with $ref
		if ref, ok := val["$ref"]; ok && !isMessage {
			refStr, ok := ref.(string)
			if !ok {
				return nil, fmt.Errorf("$ref value must be a string, got %T", ref)
			}

			refVisited := make(map[string]bool)
			for k, v := range visited {
				refVisited[k] = v
			}

			if refVisited[refStr] {
				return nil, fmt.Errorf("circular reference detected: %s", refStr)
			}
			refVisited[refStr] = true

			resolved, err := r.resolveRef(refStr)
			if err != nil {
				return nil, err
			}

			return resolved, nil
		}

		// Process regular object
		for k, v := range val {
			resolved, err := r.resolveRefsRecursive(v, visited, false)
			if err != nil {
				return nil, err
			}
			result[k] = resolved
		}
		return result, nil

	case []interface{}:
		result := make([]interface{}, len(val))
		for i, v := range val {
			resolved, err := r.resolveRefsRecursive(v, visited, false)
			if err != nil {
				return nil, err
			}
			result[i] = resolved
		}
		return result, nil

	default:
		return v, nil
	}
}

func (r *RefResolver) resolveRef(ref string) (interface{}, error) {
	if cached, ok := r.Cache[ref]; ok {
		return cached, nil
	}

	var resolved interface{}
	var err error

	switch {
	case strings.HasPrefix(ref, "#"):
		resolved, err = r.resolveLocalRef(ref)
	case strings.HasPrefix(ref, "http://") || strings.HasPrefix(ref, "https://"):
		resolved, err = r.resolveRemoteRef(ref)
	default:
		resolved, err = r.resolveFileRef(ref)
	}

	if err != nil {
		return nil, err
	}

	r.Cache[ref] = resolved
	return resolved, nil
}

func (r *RefResolver) resolveLocalRef(ref string) (interface{}, error) {
	parts := strings.Split(strings.TrimPrefix(ref, "#/"), "/")
	doc := r.Cache["#"]

	if m, ok := doc.(map[string]interface{}); ok {
		result := interface{}(m)
		for _, part := range parts {
			if m, ok := result.(map[string]interface{}); ok {
				if val, ok := m[part]; ok {
					result = val
					continue
				}
			}
			return nil, fmt.Errorf("failed to resolve path component %q in reference %s", part, ref)
		}
		return result, nil
	}
	return nil, fmt.Errorf("failed to resolve reference: %s", ref)
}

func (r *RefResolver) resolveFileRef(ref string) (interface{}, error) {
	var absPath string
	if filepath.IsAbs(ref) {
		absPath = ref
	} else if r.currentFile != "" {
		absPath = filepath.Join(filepath.Dir(r.currentFile), ref)
	} else {
		absPath = filepath.Join(r.basePath, ref)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", ref, err)
	}

	var doc interface{}
	if err := json.Unmarshal(data, &doc); err != nil {
		if err := yaml.Unmarshal(data, &doc); err != nil {
			return nil, fmt.Errorf("failed to parse file %s as YAML: %w", ref, err)
		}
	}

	return doc, nil
}

func (r *RefResolver) resolveRemoteRef(ref string) (interface{}, error) {
	resp, err := http.Get(ref)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s: %w", ref, err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from %s: %w", ref, err)
	}

	var doc interface{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("failed to parse response from %s: %w", ref, err)
	}

	return doc, nil
}