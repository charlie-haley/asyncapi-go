// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// THIS FILE IS GENERATED. DO NOT EDIT
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// If you would like to update properties for a binding,
// edit the struct for the binding you'd like to update.
// e.g sns/channel.go and run `make generate` to re-gen
// this file.
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

package sns

import (

	"encoding/json"
	"sigs.k8s.io/yaml"

)


// NewChannelBinding creates a new ChannelBinding object
func NewChannelBinding() *ChannelBinding {
	return &ChannelBinding{
	}
}


// WithName sets the 'name' field of ChannelBinding
func (obj *ChannelBinding) WithName(name string) *ChannelBinding {
	obj.Name = name
	return obj
}

// WithOrdering sets the 'ordering' field of ChannelBinding
func (obj *ChannelBinding) WithOrdering(ordering *Ordering) *ChannelBinding {
	obj.Ordering = ordering
	return obj
}

// WithPolicy sets the 'policy' field of ChannelBinding
func (obj *ChannelBinding) WithPolicy(policy *Policy) *ChannelBinding {
	obj.Policy = policy
	return obj
}

// WithTags sets the 'tags' field of ChannelBinding
func (obj *ChannelBinding) WithTags(tags map[string]string) *ChannelBinding {
	obj.Tags = tags
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



// NewOrdering creates a new Ordering object
func NewOrdering() *Ordering {
	return &Ordering{
	}
}


// WithType sets the 'type' field of Ordering
func (obj *Ordering) WithType(typeValue string) *Ordering {
	obj.Type = typeValue
	return obj
}

// WithContentBasedDeduplication sets the 'contentBasedDeduplication' field of Ordering
func (obj *Ordering) WithContentBasedDeduplication(contentBasedDeduplication bool) *Ordering {
	obj.ContentBasedDeduplication = contentBasedDeduplication
	return obj
}





// NewPolicy creates a new Policy object
func NewPolicy() *Policy {
	return &Policy{
	}
}


// WithStatements sets the 'statements' field of Policy
func (obj *Policy) WithStatements(statements []Statement) *Policy {
	obj.Statements = statements
	return obj
}





// NewOperationBinding creates a new OperationBinding object
func NewOperationBinding() *OperationBinding {
	return &OperationBinding{
	}
}


// WithTopic sets the 'topic' field of OperationBinding
func (obj *OperationBinding) WithTopic(topic *Identifier) *OperationBinding {
	obj.Topic = topic
	return obj
}

// WithConsumers sets the 'consumers' field of OperationBinding
func (obj *OperationBinding) WithConsumers(consumers []Consumer) *OperationBinding {
	obj.Consumers = consumers
	return obj
}

// WithDeliveryPolicy sets the 'deliveryPolicy' field of OperationBinding
func (obj *OperationBinding) WithDeliveryPolicy(deliveryPolicy *DeliveryPolicy) *OperationBinding {
	obj.DeliveryPolicy = deliveryPolicy
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



// NewIdentifier creates a new Identifier object
func NewIdentifier() *Identifier {
	return &Identifier{
	}
}


// WithURL sets the 'url' field of Identifier
func (obj *Identifier) WithURL(url string) *Identifier {
	obj.URL = url
	return obj
}

// WithEmail sets the 'email' field of Identifier
func (obj *Identifier) WithEmail(email string) *Identifier {
	obj.Email = email
	return obj
}

// WithPhone sets the 'phone' field of Identifier
func (obj *Identifier) WithPhone(phone string) *Identifier {
	obj.Phone = phone
	return obj
}

// WithARN sets the 'arn' field of Identifier
func (obj *Identifier) WithARN(arn string) *Identifier {
	obj.ARN = arn
	return obj
}

// WithName sets the 'name' field of Identifier
func (obj *Identifier) WithName(name string) *Identifier {
	obj.Name = name
	return obj
}





// NewDeliveryPolicy creates a new DeliveryPolicy object
func NewDeliveryPolicy() *DeliveryPolicy {
	return &DeliveryPolicy{
	}
}


// WithMinDelayTarget sets the 'minDelayTarget' field of DeliveryPolicy
func (obj *DeliveryPolicy) WithMinDelayTarget(minDelayTarget *int) *DeliveryPolicy {
	obj.MinDelayTarget = minDelayTarget
	return obj
}

// WithMaxDelayTarget sets the 'maxDelayTarget' field of DeliveryPolicy
func (obj *DeliveryPolicy) WithMaxDelayTarget(maxDelayTarget *int) *DeliveryPolicy {
	obj.MaxDelayTarget = maxDelayTarget
	return obj
}

// WithNumRetries sets the 'numRetries' field of DeliveryPolicy
func (obj *DeliveryPolicy) WithNumRetries(numRetries *int) *DeliveryPolicy {
	obj.NumRetries = numRetries
	return obj
}

// WithNumNoDelayRetries sets the 'numNoDelayRetries' field of DeliveryPolicy
func (obj *DeliveryPolicy) WithNumNoDelayRetries(numNoDelayRetries *int) *DeliveryPolicy {
	obj.NumNoDelayRetries = numNoDelayRetries
	return obj
}

// WithNumMinDelayRetries sets the 'numMinDelayRetries' field of DeliveryPolicy
func (obj *DeliveryPolicy) WithNumMinDelayRetries(numMinDelayRetries *int) *DeliveryPolicy {
	obj.NumMinDelayRetries = numMinDelayRetries
	return obj
}

// WithNumMaxDelayRetries sets the 'numMaxDelayRetries' field of DeliveryPolicy
func (obj *DeliveryPolicy) WithNumMaxDelayRetries(numMaxDelayRetries *int) *DeliveryPolicy {
	obj.NumMaxDelayRetries = numMaxDelayRetries
	return obj
}

// WithBackoffFunction sets the 'backoffFunction' field of DeliveryPolicy
func (obj *DeliveryPolicy) WithBackoffFunction(backoffFunction string) *DeliveryPolicy {
	obj.BackoffFunction = backoffFunction
	return obj
}

// WithMaxReceivesPerSecond sets the 'maxReceivesPerSecond' field of DeliveryPolicy
func (obj *DeliveryPolicy) WithMaxReceivesPerSecond(maxReceivesPerSecond *int) *DeliveryPolicy {
	obj.MaxReceivesPerSecond = maxReceivesPerSecond
	return obj
}




