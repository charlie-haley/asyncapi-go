package asyncapi2

type Components struct {
	Messages map[string]*Message `json:"messages,omitempty"`
	Schemas  map[string]*Schema  `json:"schemas,omitempty"`
	Servers  map[string]*Server  `json:"servers,omitempty"`
}

func NewComponents() *Components {
	return &Components{
		Messages: make(map[string]*Message),
		Schemas:  make(map[string]*Schema),
		Servers:  make(map[string]*Server),
	}
}

func (c *Components) WithMessage(name string, message *Message) *Components {
	c.Messages[name] = message
	return c
}

func (c *Components) WithSchema(name string, schema *Schema) *Components {
	c.Schemas[name] = schema
	return c
}

func (c *Components) WithServer(name string, server *Server) *Components {
	c.Servers[name] = server
	return c
}
