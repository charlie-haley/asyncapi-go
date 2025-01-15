package asyncapi2

type Channel struct {
	Description string                `json:"description,omitempty"`
	Parameters  map[string]*Parameter `json:"parameters,omitempty"`
	Publish     *Operation            `json:"publish,omitempty"`
	Subscribe   *Operation            `json:"subscribe,omitempty"`
	Bindings    map[string]any        `json:"bindings,omitempty"`
}

func NewChannel() *Channel {
	return &Channel{
		Parameters: make(map[string]*Parameter),
		Bindings:   make(map[string]any),
	}
}

func (c *Channel) WithDescription(description string) *Channel {
	c.Description = description
	return c
}

func (c *Channel) WithParameter(name string, parameter *Parameter) *Channel {
	c.Parameters[name] = parameter
	return c
}

func (c *Channel) WithPublish(operation *Operation) *Channel {
	c.Publish = operation
	return c
}

func (c *Channel) WithSubscribe(operation *Operation) *Channel {
	c.Subscribe = operation
	return c
}

func (c *Channel) WithBinding(name string, binding any) *Channel {
	c.Bindings[name] = binding
	return c
}
