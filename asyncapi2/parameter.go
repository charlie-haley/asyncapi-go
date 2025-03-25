package asyncapi2

type Parameter struct {
	Description string `json:"description,omitempty"`
	Schema      any    `json:"schema,omitempty"`
}

func NewParameter() *Parameter {
	return &Parameter{}
}

func (p *Parameter) WithDescription(description string) *Parameter {
	p.Description = description
	return p
}

func (p *Parameter) WithSchema(schema any) *Parameter {
	p.Schema = schema
	return p
}
