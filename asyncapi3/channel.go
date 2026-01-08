package asyncapi3

type Channel struct {
	Address     string                `json:"address,omitempty"`
	Messages    map[string]*Message   `json:"messages,omitempty"`
	Title       string                `json:"title,omitempty"`
	Summary     string                `json:"summary,omitempty"`
	Description string                `json:"description,omitempty"`
	Servers     []*Reference          `json:"servers,omitempty"`
	Parameters  map[string]*Parameter `json:"parameters,omitempty"`
	Tags        []Tag                 `json:"tags,omitempty"`
	Bindings    map[string]any        `json:"bindings,omitempty"`
}

type Parameter struct {
	Enum        []string `json:"enum,omitempty"`
	Default     string   `json:"default,omitempty"`
	Description string   `json:"description,omitempty"`
	Location    string   `json:"location,omitempty"`
}

func NewChannel() *Channel {
	return &Channel{
		Messages:   make(map[string]*Message),
		Servers:    make([]*Reference, 0),
		Parameters: make(map[string]*Parameter),
		Tags:       make([]Tag, 0),
		Bindings:   make(map[string]any),
	}
}

func (c *Channel) WithAddress(address string) *Channel {
	c.Address = address
	return c
}

func (c *Channel) WithMessage(name string, message *Message) *Channel {
	c.Messages[name] = message
	return c
}

func (c *Channel) WithTitle(title string) *Channel {
	c.Title = title
	return c
}

func (c *Channel) WithSummary(summary string) *Channel {
	c.Summary = summary
	return c
}

func (c *Channel) WithDescription(description string) *Channel {
	c.Description = description
	return c
}

func (c *Channel) WithServer(ref *Reference) *Channel {
	c.Servers = append(c.Servers, ref)
	return c
}

func (c *Channel) WithParameter(name string, param *Parameter) *Channel {
	c.Parameters[name] = param
	return c
}

func (c *Channel) WithTag(tag Tag) *Channel {
	c.Tags = append(c.Tags, tag)
	return c
}

func (c *Channel) WithBinding(name string, binding any) *Channel {
	c.Bindings[name] = binding
	return c
}

// HasBinding checks if the channel has a specific binding type
func (c *Channel) HasBinding(bindingType string) bool {
	if c.Bindings == nil {
		return false
	}
	_, ok := c.Bindings[bindingType]
	return ok
}
