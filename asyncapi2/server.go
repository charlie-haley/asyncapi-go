package asyncapi2

type Server struct {
	URL         string         `json:"url"`
	Protocol    string         `json:"protocol"`
	Description string         `json:"description,omitempty"`
	Bindings    map[string]any `json:"bindings,omitempty"`
}

func NewServer() *Server {
	return &Server{
		Bindings: make(map[string]any),
	}
}

func (s *Server) WithURL(url string) *Server {
	s.URL = url
	return s
}

func (s *Server) WithProtocol(protocol string) *Server {
	s.Protocol = protocol
	return s
}

func (s *Server) WithDescription(description string) *Server {
	s.Description = description
	return s
}

func (s *Server) WithBinding(name string, binding any) *Server {
	s.Bindings[name] = binding
	return s
}
