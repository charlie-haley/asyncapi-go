package asyncapi3

import (
	"encoding/json"
	"fmt"

	"github.com/charlie-haley/asyncapi-go/internal/validation"
)

type Document struct {
	AsyncAPI   string                `json:"asyncapi"`
	ID         string                `json:"id,omitempty"`
	Info       *Info                 `json:"info"`
	Servers    map[string]*Server    `json:"servers,omitempty"`
	Channels   map[string]*Channel   `json:"channels,omitempty"`
	Operations map[string]*Operation `json:"operations,omitempty"`
	Components *Components           `json:"components,omitempty"`
}

func NewDocument() *Document {
	return &Document{
		AsyncAPI:   "3.0.0",
		Info:       NewInfo(),
		Servers:    make(map[string]*Server),
		Channels:   make(map[string]*Channel),
		Operations: make(map[string]*Operation),
		Components: NewComponents(),
	}
}

func (d *Document) WithInfo(info *Info) *Document {
	d.Info = info
	return d
}

func (d *Document) WithServer(name string, server *Server) *Document {
	d.Servers[name] = server
	return d
}

func (d *Document) WithChannel(name string, channel *Channel) *Document {
	d.Channels[name] = channel
	return d
}

func (d *Document) WithOperation(name string, operation *Operation) *Document {
	d.Operations[name] = operation
	return d
}

func (d *Document) WithComponents(components *Components) *Document {
	d.Components = components
	return d
}

func (d *Document) Validate() error {
	if d.AsyncAPI == "" {
		return fmt.Errorf("asyncapi version is required")
	}
	if d.Info == nil {
		return fmt.Errorf("info is required")
	}

	return validation.ValidateDocument(d)
}

func (d *Document) GetVersion() string {
	return d.AsyncAPI
}

func (d *Document) MarshalJSON() ([]byte, error) {
	return json.Marshal(*d)
}

func (d *Document) UnmarshalJSON(data []byte) error {
	type Temp Document
	aux := &Temp{}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	*d = Document(*aux)
	return nil
}

// ResolveChannelRef resolves a channel reference to an actual channel
func (d *Document) ResolveChannelRef(ref string) *Channel {
	if d.Channels == nil {
		return nil
	}
	channelName := extractRefName(ref, "#/channels/")
	if channelName == "" {
		return nil
	}
	return d.Channels[channelName]
}

// ResolveMessageRef resolves a message reference to an actual message
func (d *Document) ResolveMessageRef(ref string) *Message {
	if ref == "" {
		return nil
	}

	if d.Components != nil && d.Components.Messages != nil {
		msgName := extractRefName(ref, "#/components/messages/")
		if msgName != "" {
			return d.Components.Messages[msgName]
		}
	}

	return nil
}

// GetChannelForOperation returns the channel associated with an operation
func (d *Document) GetChannelForOperation(op *Operation) *Channel {
	if op == nil || op.Channel == nil {
		return nil
	}
	return d.ResolveChannelRef(op.Channel.Ref)
}

func extractRefName(ref, prefix string) string {
	if len(ref) <= len(prefix) {
		return ""
	}
	if ref[:len(prefix)] != prefix {
		return ""
	}
	return ref[len(prefix):]
}
