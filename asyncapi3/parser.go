package asyncapi3

import "encoding/json"

func ParseFromJSON(data []byte) (*Document, error) {
	doc := NewDocument()
	if err := json.Unmarshal(data, doc); err != nil {
		return nil, err
	}
	return doc, nil
}
