package asyncapi3

type Components struct {
	Schemas           map[string]any                  `json:"schemas,omitempty"`
	Servers           map[string]*Server              `json:"servers,omitempty"`
	Channels          map[string]*Channel             `json:"channels,omitempty"`
	Operations        map[string]*Operation           `json:"operations,omitempty"`
	Messages          map[string]*Message             `json:"messages,omitempty"`
	SecuritySchemes   map[string]*SecurityScheme      `json:"securitySchemes,omitempty"`
	ServerVariables   map[string]*ServerVariable      `json:"serverVariables,omitempty"`
	Parameters        map[string]*Parameter           `json:"parameters,omitempty"`
	CorrelationIDs    map[string]*CorrelationID       `json:"correlationIds,omitempty"`
	OperationTraits   map[string]*OperationTrait      `json:"operationTraits,omitempty"`
	MessageTraits     map[string]*MessageTrait        `json:"messageTraits,omitempty"`
	ServerBindings    map[string]map[string]any       `json:"serverBindings,omitempty"`
	ChannelBindings   map[string]map[string]any       `json:"channelBindings,omitempty"`
	OperationBindings map[string]map[string]any       `json:"operationBindings,omitempty"`
	MessageBindings   map[string]map[string]any       `json:"messageBindings,omitempty"`
}

type SecurityScheme struct {
	Type             string            `json:"type"`
	Description      string            `json:"description,omitempty"`
	Name             string            `json:"name,omitempty"`
	In               string            `json:"in,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	BearerFormat     string            `json:"bearerFormat,omitempty"`
	Flows            *OAuthFlows       `json:"flows,omitempty"`
	OpenIDConnectURL string            `json:"openIdConnectUrl,omitempty"`
	Scopes           []string          `json:"scopes,omitempty"`
}

type OAuthFlows struct {
	Implicit          *OAuthFlow `json:"implicit,omitempty"`
	Password          *OAuthFlow `json:"password,omitempty"`
	ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty"`
	AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty"`
}

type OAuthFlow struct {
	AuthorizationURL string            `json:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty"`
	RefreshURL       string            `json:"refreshUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty"`
}

type OperationTrait struct {
	Title       string                `json:"title,omitempty"`
	Summary     string                `json:"summary,omitempty"`
	Description string                `json:"description,omitempty"`
	Security    []SecurityRequirement `json:"security,omitempty"`
	Tags        []Tag                 `json:"tags,omitempty"`
	Bindings    map[string]any        `json:"bindings,omitempty"`
}

type MessageTrait struct {
	Headers       any            `json:"headers,omitempty"`
	CorrelationID *CorrelationID `json:"correlationId,omitempty"`
	ContentType   string         `json:"contentType,omitempty"`
	Name          string         `json:"name,omitempty"`
	Title         string         `json:"title,omitempty"`
	Summary       string         `json:"summary,omitempty"`
	Description   string         `json:"description,omitempty"`
	Tags          []Tag          `json:"tags,omitempty"`
	Bindings      map[string]any `json:"bindings,omitempty"`
}

func NewComponents() *Components {
	return &Components{
		Schemas:           make(map[string]any),
		Servers:           make(map[string]*Server),
		Channels:          make(map[string]*Channel),
		Operations:        make(map[string]*Operation),
		Messages:          make(map[string]*Message),
		SecuritySchemes:   make(map[string]*SecurityScheme),
		ServerVariables:   make(map[string]*ServerVariable),
		Parameters:        make(map[string]*Parameter),
		CorrelationIDs:    make(map[string]*CorrelationID),
		OperationTraits:   make(map[string]*OperationTrait),
		MessageTraits:     make(map[string]*MessageTrait),
		ServerBindings:    make(map[string]map[string]any),
		ChannelBindings:   make(map[string]map[string]any),
		OperationBindings: make(map[string]map[string]any),
		MessageBindings:   make(map[string]map[string]any),
	}
}

func (c *Components) WithSchema(name string, schema any) *Components {
	c.Schemas[name] = schema
	return c
}

func (c *Components) WithServer(name string, server *Server) *Components {
	c.Servers[name] = server
	return c
}

func (c *Components) WithChannel(name string, channel *Channel) *Components {
	c.Channels[name] = channel
	return c
}

func (c *Components) WithOperation(name string, operation *Operation) *Components {
	c.Operations[name] = operation
	return c
}

func (c *Components) WithMessage(name string, message *Message) *Components {
	c.Messages[name] = message
	return c
}

func (c *Components) WithSecurityScheme(name string, scheme *SecurityScheme) *Components {
	c.SecuritySchemes[name] = scheme
	return c
}
