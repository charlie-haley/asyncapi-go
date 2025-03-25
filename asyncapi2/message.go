package asyncapi2

import "encoding/json"

type Message struct {
	Headers      any            `json:"headers,omitempty"`
	Payload      any            `json:"payload,omitempty"`
	SchemaFormat string         `json:"schemaFormat,omitempty"`
	Name         string         `json:"name,omitempty"`
	Title        string         `json:"title,omitempty"`
	Summary      string         `json:"summary,omitempty"`
	Description  string         `json:"description,omitempty"`
	ContentType  string         `json:"contentType,omitempty"`
	Bindings     map[string]any `json:"bindings,omitempty"`
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
		Bindings: make(map[string]any),
	}
}

func (m *Message) WithPayload(payload any) *Message {
	m.Payload = payload
	return m
}

func (m *Message) WithHeaders(headers any) *Message {
	m.Headers = headers
	return m
}
