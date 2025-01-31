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
	Cache    map[string]interface{}
	basePath string
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

// resolveRefsRecursive walks through the document and resolves all $ref pointers
func (r *RefResolver) resolveRefsRecursive(v interface{}, visited map[string]bool) (interface{}, error) {
	switch val := v.(type) {
	case map[string]interface{}:
		if ref, ok := val["$ref"]; ok {
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
				return nil, fmt.Errorf("failed to resolve ref %s: %w", refStr, err)
			}

			return r.resolveRefsRecursive(resolved, refVisited)
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

// resolveRef resolves a single $ref pointer
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

	// Cache the resolved reference
	r.Cache[ref] = resolved
	return resolved, nil
}

// resolveLocalRef resolves a local reference within the same document
func (r *RefResolver) resolveLocalRef(ref string) (interface{}, error) {
	parts := strings.Split(strings.TrimPrefix(ref, "#/"), "/")

	var current = r.Cache["#"]

	for _, part := range parts {
		// Unescaping JSON Pointer encoding
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

// resolveFileRef resolves a reference to a local file
func (r *RefResolver) resolveFileRef(ref string) (interface{}, error) {
	// Handle fragment
	filePath, fragment := splitFragment(ref)

	// Make path absolute if relative
	if !filepath.IsAbs(filePath) {
		filePath = filepath.Join(r.basePath, filePath)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	var doc interface{}
	if err := json.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("failed to parse file %s: %w", filePath, err)
	}

	if fragment != "" {
		return r.resolveLocalRef("#" + fragment)
	}

	return doc, nil
}

// resolveRemoteRef resolves a reference to a remote URL
func (r *RefResolver) resolveRemoteRef(ref string) (interface{}, error) {
	// Parse URL and handle fragment
	u, err := url.Parse(ref)
	if err != nil {
		return nil, fmt.Errorf("invalid URL %s: %w", ref, err)
	}

	fragment := u.Fragment
	u.Fragment = ""

	// Fetch remote document
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
		return r.resolveLocalRef("#" + fragment)
	}

	return doc, nil
}

// splitFragment splits a reference string into path and fragment parts
func splitFragment(ref string) (string, string) {
	parts := strings.SplitN(ref, "#", 2)
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}
