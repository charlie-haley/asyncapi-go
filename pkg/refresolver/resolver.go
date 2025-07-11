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
	return r.resolveRefsRecursive(doc, make(map[string]bool))
}

func (r *RefResolver) resolveRefsRecursive(v interface{}, visited map[string]bool) (interface{}, error) {
	switch val := v.(type) {
	case map[string]interface{}:
		if ref, ok := val["$ref"]; ok {
			refStr, ok := ref.(string)
			if !ok {
				return nil, fmt.Errorf("$ref value must be a string, got %T", ref)
			}

			// Check for circular refs
			if visited[refStr] {
				return nil, fmt.Errorf("circular reference detected: %s", refStr)
			}

			newVisited := copyVisitedMap(visited)
			newVisited[refStr] = true

			resolved, err := r.resolveRef(refStr)
			if err != nil {
				return nil, err
			}

			return r.resolveRefsRecursive(resolved, newVisited)
		}

		result := make(map[string]interface{})
		for k, v := range val {
			resolved, err := r.resolveRefsRecursive(v, visited)
			if err != nil {
				return nil, err
			}
			result[k] = resolved
		}
		return result, nil

	case []interface{}:
		result := make([]interface{}, len(val))
		for i, v := range val {
			resolved, err := r.resolveRefsRecursive(v, visited)
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

// Helper function to copy the visited map
func copyVisitedMap(visited map[string]bool) map[string]bool {
	newVisited := make(map[string]bool)
	for k, v := range visited {
		newVisited[k] = v
	}
	return newVisited
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
	// Handle empty fragment
	if ref == "#" {
		return r.Cache["#"], nil
	}

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

	prevFile := r.currentFile
	r.currentFile = absPath
	defer func() {
		r.currentFile = prevFile
	}()

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", absPath, err)
	}

	var doc interface{}
	if err := json.Unmarshal(data, &doc); err != nil {
		if err := yaml.Unmarshal(data, &doc); err != nil {
			return nil, fmt.Errorf("failed to parse file %s as JSON or YAML: %w", absPath, err)
		}
	}

	resolved, err := r.resolveRefsRecursive(doc, make(map[string]bool))
	if err != nil {
		return nil, fmt.Errorf("failed to resolve references in %s: %w", absPath, err)
	}

	return resolved, nil
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

	resolved, err := r.resolveRefsRecursive(doc, make(map[string]bool))
	if err != nil {
		return nil, fmt.Errorf("failed to resolve references in %s: %w", ref, err)
	}

	return resolved, nil
}
