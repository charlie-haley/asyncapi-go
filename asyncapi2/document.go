package asyncapi2

import (
	"encoding/json"
	"fmt"

	"github.com/charlie-haley/asyncapi-go/internal/validation"
)

type Document struct {
	AsyncAPI   string              `json:"asyncapi"`
	Info       *Info               `json:"info"`
	Channels   map[string]*Channel `json:"channels"`
	Servers    map[string]*Server  `json:"servers,omitempty"`
	Components *Components         `json:"components,omitempty"`
}

func NewDocument() *Document {
	return &Document{
		AsyncAPI:   "2.6.0",
		Info:       NewInfo(),
		Channels:   make(map[string]*Channel),
		Servers:    make(map[string]*Server),
		Components: NewComponents(),
	}
}

func (d *Document) WithInfo(info *Info) *Document {
	d.Info = info
	return d
}

func (d *Document) WithChannel(name string, channel *Channel) *Document {
	d.Channels[name] = channel
	return d
}

func (d *Document) WithServer(name string, server *Server) *Document {
	d.Servers[name] = server
	return d
}

func (d *Document) WithComponents(components *Components) *Document {
	d.Components = components
	return d
}

func (d *Document) Validate() error {
	// Basic validation for now
	if d.AsyncAPI == "" {
		return fmt.Errorf("asyncapi version is required")
	}
	if d.Info == nil {
		return fmt.Errorf("info is required")
	}
	if d.Channels == nil {
		return fmt.Errorf("channels is required")
	}

	// Schema validation
	return validation.ValidateDocument(d)
}

// GetVersion implements spec.Document.
func (d *Document) GetVersion() string {
	return d.AsyncAPI
}

// MarshalJSON implements spec.Document.
func (d *Document) MarshalJSON() ([]byte, error) {
	return json.Marshal(*d)
}

// UnmarshalJSON implements spec.Document.
func (d *Document) UnmarshalJSON(data []byte) error {
	// Create an temp type to prevent infinite recursion
	type Temp Document
	aux := &Temp{}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	// Copy the data from the temp type to the main struct
	*d = Document(*aux)

	return nil
}
