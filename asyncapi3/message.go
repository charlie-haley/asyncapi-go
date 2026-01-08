package asyncapi3

import "encoding/json"

type Message struct {
	Headers       any            `json:"headers,omitempty"`
	Payload       any            `json:"payload,omitempty"`
	CorrelationID *CorrelationID `json:"correlationId,omitempty"`
	ContentType   string         `json:"contentType,omitempty"`
	Name          string         `json:"name,omitempty"`
	Title         string         `json:"title,omitempty"`
	Summary       string         `json:"summary,omitempty"`
	Description   string         `json:"description,omitempty"`
	Tags          []Tag          `json:"tags,omitempty"`
	Bindings      map[string]any `json:"bindings,omitempty"`
	Traits        []*Reference   `json:"traits,omitempty"`
}

type CorrelationID struct {
	Description string `json:"description,omitempty"`
	Location    string `json:"location"`
}

func (m *Message) UnmarshalJSON(data []byte) error {
	type MessageAlias Message
	temp := &MessageAlias{}
	if err := json.Unmarshal(data, temp); err != nil {
		return err
	}
	*m = Message(*temp)
	return nil
}

func NewMessage() *Message {
	return &Message{
		Tags:     make([]Tag, 0),
		Bindings: make(map[string]any),
		Traits:   make([]*Reference, 0),
	}
}

func (m *Message) WithHeaders(headers any) *Message {
	m.Headers = headers
	return m
}

func (m *Message) WithPayload(payload any) *Message {
	m.Payload = payload
	return m
}

func (m *Message) WithContentType(contentType string) *Message {
	m.ContentType = contentType
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

func (m *Message) WithTag(tag Tag) *Message {
	m.Tags = append(m.Tags, tag)
	return m
}

func (m *Message) WithBinding(name string, binding any) *Message {
	m.Bindings[name] = binding
	return m
}
