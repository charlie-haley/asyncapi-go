// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// THIS FILE IS GENERATED. DO NOT EDIT
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// If you would like to update properties for a binding,
// edit the struct for the binding you'd like to update.
// e.g sqs/channel.go and run `make generate` to re-gen
// this file.
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

package sqs

import (

	"encoding/json"
	"sigs.k8s.io/yaml"

)


// NewOperationBinding creates a new OperationBinding object
func NewOperationBinding() *OperationBinding {
	return &OperationBinding{
	}
}


// WithQueues sets the 'queues' field of OperationBinding
func (obj *OperationBinding) WithQueues(queues []Queue) *OperationBinding {
	obj.Queues = queues
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



// NewChannelBinding creates a new ChannelBinding object
func NewChannelBinding() *ChannelBinding {
	return &ChannelBinding{
		BindingVersion: "latest",
	}
}


// WithQueue sets the 'queue' field of ChannelBinding
func (obj *ChannelBinding) WithQueue(queue *Queue) *ChannelBinding {
	obj.Queue = queue
	return obj
}

// WithDeadLetterQueue sets the 'deadLetterQueue' field of ChannelBinding
func (obj *ChannelBinding) WithDeadLetterQueue(deadLetterQueue *Queue) *ChannelBinding {
	obj.DeadLetterQueue = deadLetterQueue
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



// NewQueue creates a new Queue object
func NewQueue() *Queue {
	return &Queue{
	}
}


// WithName sets the 'name' field of Queue
func (obj *Queue) WithName(name string) *Queue {
	obj.Name = name
	return obj
}

// WithFifoQueue sets the 'fifoQueue' field of Queue
func (obj *Queue) WithFifoQueue(fifoQueue bool) *Queue {
	obj.FifoQueue = fifoQueue
	return obj
}

// WithDeduplicationScope sets the 'deduplicationScope' field of Queue
func (obj *Queue) WithDeduplicationScope(deduplicationScope string) *Queue {
	obj.DeduplicationScope = deduplicationScope
	return obj
}

// WithFifoThroughputLimit sets the 'fifoThroughputLimit' field of Queue
func (obj *Queue) WithFifoThroughputLimit(fifoThroughputLimit string) *Queue {
	obj.FifoThroughputLimit = fifoThroughputLimit
	return obj
}

// WithDeliveryDelay sets the 'deliveryDelay' field of Queue
func (obj *Queue) WithDeliveryDelay(deliveryDelay int) *Queue {
	obj.DeliveryDelay = deliveryDelay
	return obj
}

// WithVisibilityTimeout sets the 'visibilityTimeout' field of Queue
func (obj *Queue) WithVisibilityTimeout(visibilityTimeout int) *Queue {
	obj.VisibilityTimeout = visibilityTimeout
	return obj
}

// WithReceiveMessageWaitTime sets the 'receiveMessageWaitTime' field of Queue
func (obj *Queue) WithReceiveMessageWaitTime(receiveMessageWaitTime int) *Queue {
	obj.ReceiveMessageWaitTime = receiveMessageWaitTime
	return obj
}

// WithMessageRetentionPeriod sets the 'messageRetentionPeriod' field of Queue
func (obj *Queue) WithMessageRetentionPeriod(messageRetentionPeriod int) *Queue {
	obj.MessageRetentionPeriod = messageRetentionPeriod
	return obj
}

// WithRedrivePolicy sets the 'redrivePolicy' field of Queue
func (obj *Queue) WithRedrivePolicy(redrivePolicy *RedrivePolicy) *Queue {
	obj.RedrivePolicy = redrivePolicy
	return obj
}

// WithPolicy sets the 'policy' field of Queue
func (obj *Queue) WithPolicy(policy *Policy) *Queue {
	obj.Policy = policy
	return obj
}

// WithTags sets the 'tags' field of Queue
func (obj *Queue) WithTags(tags map[string]string) *Queue {
	obj.Tags = tags
	return obj
}




