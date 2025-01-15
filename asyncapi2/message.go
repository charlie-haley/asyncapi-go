package asyncapi2

type Message struct {
	Headers     *Schema        `json:"headers,omitempty"`
	Payload     any            `json:"payload,omitempty"`
	Name        string         `json:"name,omitempty"`
	Title       string         `json:"title,omitempty"`
	Summary     string         `json:"summary,omitempty"`
	Description string         `json:"description,omitempty"`
	ContentType string         `json:"contentType,omitempty"`
	Bindings    map[string]any `json:"bindings,omitempty"`
}

func NewMessage() *Message {
	return &Message{
		Bindings: make(map[string]any),
	}
}

func (m *Message) WithHeaders(headers *Schema) *Message {
	m.Headers = headers
	return m
}

func (m *Message) WithPayload(payload any) *Message {
	m.Payload = payload
	return m
}

func (m *Message) WithName(name string) *Message {
	m.Name = name
	return m
}

func (m *Message) WithTitle(title string) *Message {
	m.Title = title
	return m
}

func (m *Message) WithSummary(summary string) *Message {
	m.Summary = summary
	return m
}

func (m *Message) WithDescription(description string) *Message {
	m.Description = description
	return m
}

func (m *Message) WithContentType(contentType string) *Message {
	m.ContentType = contentType
	return m
}

func (m *Message) WithBinding(name string, binding any) *Message {
	m.Bindings[name] = binding
	return m
}
