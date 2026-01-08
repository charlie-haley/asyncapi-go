package asyncapi3

import "github.com/charlie-haley/asyncapi-go/spec"

type Operation struct {
	Action      spec.Action    `json:"action"`
	Channel     *Reference     `json:"channel"`
	Title       string         `json:"title,omitempty"`
	Summary     string         `json:"summary,omitempty"`
	Description string         `json:"description,omitempty"`
	Security    []SecurityRequirement `json:"security,omitempty"`
	Tags        []Tag                 `json:"tags,omitempty"`
	Messages    []*Reference          `json:"messages,omitempty"`
	Reply       *OperationReply       `json:"reply,omitempty"`
	Bindings    map[string]any        `json:"bindings,omitempty"`
}

type OperationReply struct {
	Address  *OperationReplyAddress `json:"address,omitempty"`
	Channel  *Reference             `json:"channel,omitempty"`
	Messages []*Reference           `json:"messages,omitempty"`
}

type OperationReplyAddress struct {
	Location    string `json:"location"`
	Description string `json:"description,omitempty"`
}

func NewOperation() *Operation {
	return &Operation{
		Security: make([]SecurityRequirement, 0),
		Tags:     make([]Tag, 0),
		Messages: make([]*Reference, 0),
		Bindings: make(map[string]any),
	}
}

func (o *Operation) WithAction(action spec.Action) *Operation {
	o.Action = action
	return o
}

func (o *Operation) WithChannel(ref *Reference) *Operation {
	o.Channel = ref
	return o
}

func (o *Operation) WithTitle(title string) *Operation {
	o.Title = title
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

func (o *Operation) WithMessage(ref *Reference) *Operation {
	o.Messages = append(o.Messages, ref)
	return o
}

func (o *Operation) WithTag(tag Tag) *Operation {
	o.Tags = append(o.Tags, tag)
	return o
}

func (o *Operation) WithBinding(name string, binding any) *Operation {
	o.Bindings[name] = binding
	return o
}

// IsSend returns true if this operation sends messages
func (o *Operation) IsSend() bool {
	return o.Action == spec.Send
}

// IsReceive returns true if this operation receives messages
func (o *Operation) IsReceive() bool {
	return o.Action == spec.Receive
}
