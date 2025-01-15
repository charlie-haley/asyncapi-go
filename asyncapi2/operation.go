package asyncapi2

type Operation struct {
	OperationID string         `json:"operationId,omitempty"`
	Summary     string         `json:"summary,omitempty"`
	Description string         `json:"description,omitempty"`
	Tags        []Tag          `json:"tags,omitempty"`
	Message     *Message       `json:"message,omitempty"`
	Bindings    map[string]any `json:"bindings,omitempty"`
}

func NewOperation() *Operation {
	return &Operation{
		Tags:     make([]Tag, 0),
		Bindings: make(map[string]any),
	}
}

func (o *Operation) WithOperationID(id string) *Operation {
	o.OperationID = id
	return o
}

func (o *Operation) WithSummary(summary string) *Operation {
	o.Summary = summary
	return o
}

func (o *Operation) WithDescription(description string) *Operation {
	o.Description = description
	return o
}

func (o *Operation) WithTag(tag Tag) *Operation {
	o.Tags = append(o.Tags, tag)
	return o
}

func (o *Operation) WithMessage(message *Message) *Operation {
	o.Message = message
	return o
}

func (o *Operation) WithBinding(name string, binding any) *Operation {
	o.Bindings[name] = binding
	return o
}
