package refresolver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// RefResolver handles resolving references in AsyncAPI documents
type RefResolver struct {
	Cache       map[string]interface{}
	basePath    string
	currentFile string
}

// New creates a new RefResolver instance
func New(basePath string) *RefResolver {
	return &RefResolver{
		Cache:    make(map[string]interface{}),
		basePath: basePath,
	}
}

// ResolveRefs takes a JSON/YAML document and resolves all $ref pointers
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

			// Create a new visited map for this reference path
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

			previousFile := r.currentFile

			if !strings.HasPrefix(refStr, "#") && !strings.HasPrefix(refStr, "http") {
				if filepath.IsAbs(refStr) {
					r.currentFile = refStr
				} else {
					// If we have a current file, resolve relative to it
					if r.currentFile != "" {
						r.currentFile = filepath.Join(filepath.Dir(r.currentFile), refStr)
					} else {
						r.currentFile = filepath.Join(r.basePath, refStr)
					}
				}
			}

			result, err := r.resolveRefsRecursive(resolved, refVisited)

			r.currentFile = previousFile

			return result, err
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

	var current = r.Cache["#"]

	for _, part := range parts {
		part = strings.ReplaceAll(part, "~1", "/")
		part = strings.ReplaceAll(part, "~0", "~")

		switch v := current.(type) {
		case map[string]interface{}:
			var ok bool
			current, ok = v[part]
			if !ok {
				return nil, fmt.Errorf("failed to resolve reference: %s not found", part)
			}
		default:
			return nil, fmt.Errorf("invalid reference path: %s is not an object", part)
		}
	}

	return current, nil
}

func (r *RefResolver) resolveFileRef(ref string) (interface{}, error) {
	filePath, fragment := splitFragment(ref)

	var absPath string
	if filepath.IsAbs(filePath) {
		absPath = filePath
	} else if r.currentFile != "" {
		// If we have a current file, resolve relative to it
		absPath = filepath.Join(filepath.Dir(r.currentFile), filePath)
	} else {
		// Otherwise resolve relative to base path
		absPath = filepath.Join(r.basePath, filePath)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	var doc interface{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("failed to parse file %s: %w", filePath, err)
	}

	if fragment != "" {
		r.Cache["#"] = doc
		return r.resolveLocalRef("#" + fragment)
	}

	return doc, nil
}

func (r *RefResolver) resolveRemoteRef(ref string) (interface{}, error) {
	u, err := url.Parse(ref)
	if err != nil {
		return nil, fmt.Errorf("invalid URL %s: %w", ref, err)
	}

	fragment := u.Fragment
	u.Fragment = ""

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s: %w", u.String(), err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from %s: %w", u.String(), err)
	}

	var doc interface{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("failed to parse response from %s: %w", u.String(), err)
	}

	if fragment != "" {
		r.Cache["#"] = doc
		return r.resolveLocalRef("#" + fragment)
	}

	return doc, nil
}

func splitFragment(ref string) (string, string) {
	parts := strings.SplitN(ref, "#", 2)
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}
