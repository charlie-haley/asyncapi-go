package amqp

// ChannelBinding represents the AMQP 0-9-1 Channel Binding object.
//
// This object contains information about the channel representation in AMQP.
// +binding
type ChannelBinding struct {
	Is             string    `json:"is"`
	Exchange       *Exchange `json:"exchange,omitempty"`
	Queue          *Queue    `json:"queue,omitempty"`
	BindingVersion string    `json:"bindingVersion,omitempty"`
}

// Exchange is the object that defines the exchange properties when the channel is a routing key.
type Exchange struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"autoDelete"`
    VHost      string `json:"vhost,omitempty" default:"\"/\""`
}

// ExchangeType represents the different types of exchanges.
const (
	ExchangeTypeTopic   = "topic"
	ExchangeTypeDirect  = "direct"
	ExchangeTypeFanout  = "fanout"
	ExchangeTypeDefault = "default"
	ExchangeTypeHeaders = "headers"
)

// Queue is the object that defines the queue properties when the channel is a queue.
type Queue struct {
	Name       string `json:"name"`
	Durable    bool   `json:"durable"`
	Exclusive  bool   `json:"exclusive"`
	AutoDelete bool   `json:"autoDelete"`
    VHost      string `json:"vhost,omitempty" default:"\"/\""`
}

// ChannelIs represents the different types of channels.
const (
	ChannelIsRoutingKey = "routingKey"
	ChannelIsQueue      = "queue"
)
