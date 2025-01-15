// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// THIS FILE IS GENERATED. DO NOT EDIT
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// If you would like to update properties for a binding,
// edit the struct for the binding you'd like to update.
// e.g amqp/channel.go and run `make generate` to re-gen
// this file.
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

package amqp

import (

	"encoding/json"
	"sigs.k8s.io/yaml"

)


// NewChannelBinding creates a new ChannelBinding object
func NewChannelBinding() *ChannelBinding {
	return &ChannelBinding{
	}
}


// WithIs sets the 'is' field of ChannelBinding
func (obj *ChannelBinding) WithIs(is string) *ChannelBinding {
	obj.Is = is
	return obj
}

// WithExchange sets the 'exchange' field of ChannelBinding
func (obj *ChannelBinding) WithExchange(exchange *Exchange) *ChannelBinding {
	obj.Exchange = exchange
	return obj
}

// WithQueue sets the 'queue' field of ChannelBinding
func (obj *ChannelBinding) WithQueue(queue *Queue) *ChannelBinding {
	obj.Queue = queue
	return obj
}

// WithBindingVersion sets the 'bindingVersion' field of ChannelBinding
func (obj *ChannelBinding) WithBindingVersion(bindingVersion string) *ChannelBinding {
	obj.BindingVersion = bindingVersion
	return obj
}



// MarshalYAML is a custom marshaller that converts ChannelBinding to YAML
func (t ChannelBinding) MarshalYAML() (interface{}, error) {
    bytes, err := json.Marshal(t)
    if err != nil {
        return nil, err
    }
    var out interface{}
    err = yaml.Unmarshal(bytes, &out) 
    return out, err
}

// UnmarshalYAML is a custom unmarshaler that converts YAML to ChannelBinding
func (t *ChannelBinding) UnmarshalYAML(unmarshal func(interface{}) error) error {
    var temp interface{}
    if err := unmarshal(&temp); err != nil {
        return err
    }
    bytes, err := yaml.Marshal(temp)
    if err != nil {
        return err
    }
    return json.Unmarshal(bytes, t)
}

// MarshalJSON is a custom marshaller that converts ChannelBinding to JSON
func (t ChannelBinding) MarshalJSON() ([]byte, error) {
	type Alias ChannelBinding
	return json.Marshal(struct{ Alias }{Alias(t)})
}

// UnmarshalJSON is a custom unmarshaler that converts JSON to ChannelBinding
func (t *ChannelBinding) UnmarshalJSON(data []byte) error {
	type Alias ChannelBinding
	aux := struct{ *Alias }{Alias: (*Alias)(t)}
	return json.Unmarshal(data, &aux)
}



// NewExchange creates a new Exchange object
func NewExchange() *Exchange {
	return &Exchange{
		VHost: "/",
	}
}


// WithName sets the 'name' field of Exchange
func (obj *Exchange) WithName(name string) *Exchange {
	obj.Name = name
	return obj
}

// WithType sets the 'type' field of Exchange
func (obj *Exchange) WithType(typeValue string) *Exchange {
	obj.Type = typeValue
	return obj
}

// WithDurable sets the 'durable' field of Exchange
func (obj *Exchange) WithDurable(durable bool) *Exchange {
	obj.Durable = durable
	return obj
}

// WithAutoDelete sets the 'autoDelete' field of Exchange
func (obj *Exchange) WithAutoDelete(autoDelete bool) *Exchange {
	obj.AutoDelete = autoDelete
	return obj
}

// WithVHost sets the 'vhost' field of Exchange
func (obj *Exchange) WithVHost(vhost string) *Exchange {
	obj.VHost = vhost
	return obj
}





// NewQueue creates a new Queue object
func NewQueue() *Queue {
	return &Queue{
		VHost: "/",
	}
}


// WithName sets the 'name' field of Queue
func (obj *Queue) WithName(name string) *Queue {
	obj.Name = name
	return obj
}

// WithDurable sets the 'durable' field of Queue
func (obj *Queue) WithDurable(durable bool) *Queue {
	obj.Durable = durable
	return obj
}

// WithExclusive sets the 'exclusive' field of Queue
func (obj *Queue) WithExclusive(exclusive bool) *Queue {
	obj.Exclusive = exclusive
	return obj
}

// WithAutoDelete sets the 'autoDelete' field of Queue
func (obj *Queue) WithAutoDelete(autoDelete bool) *Queue {
	obj.AutoDelete = autoDelete
	return obj
}

// WithVHost sets the 'vhost' field of Queue
func (obj *Queue) WithVHost(vhost string) *Queue {
	obj.VHost = vhost
	return obj
}





// NewMessageBinding creates a new MessageBinding object
func NewMessageBinding() *MessageBinding {
	return &MessageBinding{
	}
}


// WithContentEncoding sets the 'contentEncoding' field of MessageBinding
func (obj *MessageBinding) WithContentEncoding(contentEncoding string) *MessageBinding {
	obj.ContentEncoding = contentEncoding
	return obj
}

// WithMessageType sets the 'messageType' field of MessageBinding
func (obj *MessageBinding) WithMessageType(messageType string) *MessageBinding {
	obj.MessageType = messageType
	return obj
}

// WithBindingVersion sets the 'bindingVersion' field of MessageBinding
func (obj *MessageBinding) WithBindingVersion(bindingVersion string) *MessageBinding {
	obj.BindingVersion = bindingVersion
	return obj
}



// MarshalYAML is a custom marshaller that converts MessageBinding to YAML
func (t MessageBinding) MarshalYAML() (interface{}, error) {
    bytes, err := json.Marshal(t)
    if err != nil {
        return nil, err
    }
    var out interface{}
    err = yaml.Unmarshal(bytes, &out) 
    return out, err
}

// UnmarshalYAML is a custom unmarshaler that converts YAML to MessageBinding
func (t *MessageBinding) UnmarshalYAML(unmarshal func(interface{}) error) error {
    var temp interface{}
    if err := unmarshal(&temp); err != nil {
        return err
    }
    bytes, err := yaml.Marshal(temp)
    if err != nil {
        return err
    }
    return json.Unmarshal(bytes, t)
}

// MarshalJSON is a custom marshaller that converts MessageBinding to JSON
func (t MessageBinding) MarshalJSON() ([]byte, error) {
	type Alias MessageBinding
	return json.Marshal(struct{ Alias }{Alias(t)})
}

// UnmarshalJSON is a custom unmarshaler that converts JSON to MessageBinding
func (t *MessageBinding) UnmarshalJSON(data []byte) error {
	type Alias MessageBinding
	aux := struct{ *Alias }{Alias: (*Alias)(t)}
	return json.Unmarshal(data, &aux)
}



// NewOperationBinding creates a new OperationBinding object
func NewOperationBinding() *OperationBinding {
	return &OperationBinding{
	}
}


// WithExpiration sets the 'expiration' field of OperationBinding
func (obj *OperationBinding) WithExpiration(expiration int) *OperationBinding {
	obj.Expiration = expiration
	return obj
}

// WithUserID sets the 'userId' field of OperationBinding
func (obj *OperationBinding) WithUserID(userId string) *OperationBinding {
	obj.UserID = userId
	return obj
}

// WithCC sets the 'cc' field of OperationBinding
func (obj *OperationBinding) WithCC(cc []string) *OperationBinding {
	obj.CC = cc
	return obj
}

// WithPriority sets the 'priority' field of OperationBinding
func (obj *OperationBinding) WithPriority(priority int) *OperationBinding {
	obj.Priority = priority
	return obj
}

// WithDeliveryMode sets the 'deliveryMode' field of OperationBinding
func (obj *OperationBinding) WithDeliveryMode(deliveryMode int) *OperationBinding {
	obj.DeliveryMode = deliveryMode
	return obj
}

// WithMandatory sets the 'mandatory' field of OperationBinding
func (obj *OperationBinding) WithMandatory(mandatory bool) *OperationBinding {
	obj.Mandatory = mandatory
	return obj
}

// WithBCC sets the 'bcc' field of OperationBinding
func (obj *OperationBinding) WithBCC(bcc []string) *OperationBinding {
	obj.BCC = bcc
	return obj
}

// WithTimestamp sets the 'timestamp' field of OperationBinding
func (obj *OperationBinding) WithTimestamp(timestamp bool) *OperationBinding {
	obj.Timestamp = timestamp
	return obj
}

// WithAck sets the 'ack' field of OperationBinding
func (obj *OperationBinding) WithAck(ack bool) *OperationBinding {
	obj.Ack = ack
	return obj
}

// WithBindingVersion sets the 'bindingVersion' field of OperationBinding
func (obj *OperationBinding) WithBindingVersion(bindingVersion string) *OperationBinding {
	obj.BindingVersion = bindingVersion
	return obj
}



// MarshalYAML is a custom marshaller that converts OperationBinding to YAML
func (t OperationBinding) MarshalYAML() (interface{}, error) {
    bytes, err := json.Marshal(t)
    if err != nil {
        return nil, err
    }
    var out interface{}
    err = yaml.Unmarshal(bytes, &out) 
    return out, err
}

// UnmarshalYAML is a custom unmarshaler that converts YAML to OperationBinding
func (t *OperationBinding) UnmarshalYAML(unmarshal func(interface{}) error) error {
    var temp interface{}
    if err := unmarshal(&temp); err != nil {
        return err
    }
    bytes, err := yaml.Marshal(temp)
    if err != nil {
        return err
    }
    return json.Unmarshal(bytes, t)
}

// MarshalJSON is a custom marshaller that converts OperationBinding to JSON
func (t OperationBinding) MarshalJSON() ([]byte, error) {
	type Alias OperationBinding
	return json.Marshal(struct{ Alias }{Alias(t)})
}

// UnmarshalJSON is a custom unmarshaler that converts JSON to OperationBinding
func (t *OperationBinding) UnmarshalJSON(data []byte) error {
	type Alias OperationBinding
	aux := struct{ *Alias }{Alias: (*Alias)(t)}
	return json.Unmarshal(data, &aux)
}


