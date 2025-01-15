package asyncapi2

type Schema struct {
	Type                 string             `json:"type,omitempty"`
	Items                *Schema            `json:"items,omitempty"`
	Properties           map[string]*Schema `json:"properties,omitempty"`
	Required             []string           `json:"required,omitempty"`
	AdditionalProperties any                `json:"additionalProperties,omitempty"`
}

func NewSchema() *Schema {
	return &Schema{
		Properties: make(map[string]*Schema),
		Required:   make([]string, 0),
	}
}

func (s *Schema) WithType(type_ string) *Schema {
	s.Type = type_
	return s
}

func (s *Schema) WithItems(items *Schema) *Schema {
	s.Items = items
	return s
}

func (s *Schema) WithProperty(name string, schema *Schema) *Schema {
	s.Properties[name] = schema
	return s
}

func (s *Schema) WithRequired(required []string) *Schema {
	s.Required = required
	return s
}

func (s *Schema) WithAdditionalProperties(additionalProperties any) *Schema {
	s.AdditionalProperties = additionalProperties
	return s
}
