package asyncapi3

type Server struct {
	Host        string            `json:"host"`
	Protocol    string            `json:"protocol"`
	ProtocolVer string            `json:"protocolVersion,omitempty"`
	Pathname    string            `json:"pathname,omitempty"`
	Description string            `json:"description,omitempty"`
	Title       string            `json:"title,omitempty"`
	Summary     string            `json:"summary,omitempty"`
	Variables   map[string]*ServerVariable `json:"variables,omitempty"`
	Security    []SecurityRequirement      `json:"security,omitempty"`
	Tags        []Tag                      `json:"tags,omitempty"`
	Bindings    map[string]any             `json:"bindings,omitempty"`
}

type ServerVariable struct {
	Enum        []string `json:"enum,omitempty"`
	Default     string   `json:"default,omitempty"`
	Description string   `json:"description,omitempty"`
}

type SecurityRequirement map[string][]string

func NewServer() *Server {
	return &Server{
		Variables: make(map[string]*ServerVariable),
		Security:  make([]SecurityRequirement, 0),
		Tags:      make([]Tag, 0),
		Bindings:  make(map[string]any),
	}
}

func (s *Server) WithHost(host string) *Server {
	s.Host = host
	return s
}

func (s *Server) WithProtocol(protocol string) *Server {
	s.Protocol = protocol
	return s
}

func (s *Server) WithProtocolVersion(version string) *Server {
	s.ProtocolVer = version
	return s
}

func (s *Server) WithPathname(pathname string) *Server {
	s.Pathname = pathname
	return s
}

func (s *Server) WithDescription(description string) *Server {
	s.Description = description
	return s
}

func (s *Server) WithVariable(name string, variable *ServerVariable) *Server {
	s.Variables[name] = variable
	return s
}

func (s *Server) WithBinding(name string, binding any) *Server {
	s.Bindings[name] = binding
	return s
}
