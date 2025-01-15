package spec

// Action specifies if an operation sends or receives messages
type Action string

const (
	Send    Action = "send"
	Receive Action = "receive"
)

// Document represents the core interface that all AsyncAPI versions implement
type Document interface {
	Validate() error
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
	GetVersion() string
}

// BaseInfo provides a common implementation of the Info interface
type BaseInfo struct {
	Title       string `json:"title"`
	Version     string `json:"version"`
	Description string `json:"description,omitempty"`
}

// BaseDocument contains fields common to all AsyncAPI versions
type BaseDocument struct {
	Version string   `json:"asyncapi"`
	Info    BaseInfo `json:"info"`
}

func (d *BaseDocument) GetVersion() string {
	return d.Version
}