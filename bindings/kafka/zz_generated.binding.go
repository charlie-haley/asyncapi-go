// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// THIS FILE IS GENERATED. DO NOT EDIT
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// If you would like to update properties for a binding,
// edit the struct for the binding you'd like to update.
// e.g kafka/channel.go and run `make generate` to re-gen
// this file.
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

package kafka

import (

	"encoding/json"
	"sigs.k8s.io/yaml"

)


// NewMessageBinding creates a new MessageBinding object
func NewMessageBinding() *MessageBinding {
	return &MessageBinding{
	}
}


// WithKey sets the 'key' field of MessageBinding
func (obj *MessageBinding) WithKey(key interface{}) *MessageBinding {
	obj.Key = key
	return obj
}

// WithSchemaIDLocation sets the 'schemaIdLocation' field of MessageBinding
func (obj *MessageBinding) WithSchemaIDLocation(schemaIdLocation string) *MessageBinding {
	obj.SchemaIDLocation = schemaIdLocation
	return obj
}

// WithSchemaIDPayloadEncoding sets the 'schemaIdPayloadEncoding' field of MessageBinding
func (obj *MessageBinding) WithSchemaIDPayloadEncoding(schemaIdPayloadEncoding string) *MessageBinding {
	obj.SchemaIDPayloadEncoding = schemaIdPayloadEncoding
	return obj
}

// WithSchemaLookupStrategy sets the 'schemaLookupStrategy' field of MessageBinding
func (obj *MessageBinding) WithSchemaLookupStrategy(schemaLookupStrategy string) *MessageBinding {
	obj.SchemaLookupStrategy = schemaLookupStrategy
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



// NewChannelBinding creates a new ChannelBinding object
func NewChannelBinding() *ChannelBinding {
	return &ChannelBinding{
	}
}


// WithTopic sets the 'topic' field of ChannelBinding
func (obj *ChannelBinding) WithTopic(topic string) *ChannelBinding {
	obj.Topic = topic
	return obj
}

// WithPartitions sets the 'partitions' field of ChannelBinding
func (obj *ChannelBinding) WithPartitions(partitions int) *ChannelBinding {
	obj.Partitions = partitions
	return obj
}

// WithReplicas sets the 'replicas' field of ChannelBinding
func (obj *ChannelBinding) WithReplicas(replicas int) *ChannelBinding {
	obj.Replicas = replicas
	return obj
}

// WithTopicConfiguration sets the 'topicConfiguration' field of ChannelBinding
func (obj *ChannelBinding) WithTopicConfiguration(topicConfiguration *TopicConfiguration) *ChannelBinding {
	obj.TopicConfiguration = topicConfiguration
	return obj
}

// WithBindingVersion sets the 'bindingVersion' field of ChannelBinding
func (obj *ChannelBinding) WithBindingVersion(bindingVersion string) *ChannelBinding {
	obj.BindingVersion = bindingVersion
	return obj
}





// NewTopicConfiguration creates a new TopicConfiguration object
func NewTopicConfiguration() *TopicConfiguration {
	return &TopicConfiguration{
	}
}


// WithCleanupPolicy sets the 'cleanup.policy' field of TopicConfiguration
func (obj *TopicConfiguration) WithCleanupPolicy(cleanuppolicy []string) *TopicConfiguration {
	obj.CleanupPolicy = cleanuppolicy
	return obj
}

// WithRetentionMs sets the 'retention.ms' field of TopicConfiguration
func (obj *TopicConfiguration) WithRetentionMs(retentionms int64) *TopicConfiguration {
	obj.RetentionMs = retentionms
	return obj
}

// WithRetentionBytes sets the 'retention.bytes' field of TopicConfiguration
func (obj *TopicConfiguration) WithRetentionBytes(retentionbytes int64) *TopicConfiguration {
	obj.RetentionBytes = retentionbytes
	return obj
}

// WithDeleteRetentionMs sets the 'delete.retention.ms' field of TopicConfiguration
func (obj *TopicConfiguration) WithDeleteRetentionMs(deleteretentionms int64) *TopicConfiguration {
	obj.DeleteRetentionMs = deleteretentionms
	return obj
}

// WithMaxMessageBytes sets the 'max.message.bytes' field of TopicConfiguration
func (obj *TopicConfiguration) WithMaxMessageBytes(maxmessagebytes int) *TopicConfiguration {
	obj.MaxMessageBytes = maxmessagebytes
	return obj
}

// WithConfluentKeySchemaValidation sets the 'confluent.key.schema.validation' field of TopicConfiguration
func (obj *TopicConfiguration) WithConfluentKeySchemaValidation(confluentkeyschemavalidation bool) *TopicConfiguration {
	obj.ConfluentKeySchemaValidation = confluentkeyschemavalidation
	return obj
}

// WithConfluentKeySubjectNameStrategy sets the 'confluent.key.subject.name.strategy' field of TopicConfiguration
func (obj *TopicConfiguration) WithConfluentKeySubjectNameStrategy(confluentkeysubjectnamestrategy string) *TopicConfiguration {
	obj.ConfluentKeySubjectNameStrategy = confluentkeysubjectnamestrategy
	return obj
}

// WithConfluentValueSchemaValidation sets the 'confluent.value.schema.validation' field of TopicConfiguration
func (obj *TopicConfiguration) WithConfluentValueSchemaValidation(confluentvalueschemavalidation bool) *TopicConfiguration {
	obj.ConfluentValueSchemaValidation = confluentvalueschemavalidation
	return obj
}

// WithConfluentValueSubjectNameStrategy sets the 'confluent.value.subject.name.strategy' field of TopicConfiguration
func (obj *TopicConfiguration) WithConfluentValueSubjectNameStrategy(confluentvaluesubjectnamestrategy string) *TopicConfiguration {
	obj.ConfluentValueSubjectNameStrategy = confluentvaluesubjectnamestrategy
	return obj
}

// WithAdditionalProperties sets the '-' field of TopicConfiguration
func (obj *TopicConfiguration) WithAdditionalProperties(additionalProperties map[string]interface{}) *TopicConfiguration {
	obj.AdditionalProperties = additionalProperties
	return obj
}





// NewOperationBinding creates a new OperationBinding object
func NewOperationBinding() *OperationBinding {
	return &OperationBinding{
	}
}


// WithGroupID sets the 'groupId' field of OperationBinding
func (obj *OperationBinding) WithGroupID(groupId interface{}) *OperationBinding {
	obj.GroupID = groupId
	return obj
}

// WithClientID sets the 'clientId' field of OperationBinding
func (obj *OperationBinding) WithClientID(clientId interface{}) *OperationBinding {
	obj.ClientID = clientId
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



// NewServerBinding creates a new ServerBinding object
func NewServerBinding() *ServerBinding {
	return &ServerBinding{
	}
}


// WithSchemaRegistryURL sets the 'schemaRegistryUrl' field of ServerBinding
func (obj *ServerBinding) WithSchemaRegistryURL(schemaRegistryUrl string) *ServerBinding {
	obj.SchemaRegistryURL = schemaRegistryUrl
	return obj
}

// WithSchemaRegistryVendor sets the 'schemaRegistryVendor' field of ServerBinding
func (obj *ServerBinding) WithSchemaRegistryVendor(schemaRegistryVendor string) *ServerBinding {
	obj.SchemaRegistryVendor = schemaRegistryVendor
	return obj
}

// WithBindingVersion sets the 'bindingVersion' field of ServerBinding
func (obj *ServerBinding) WithBindingVersion(bindingVersion string) *ServerBinding {
	obj.BindingVersion = bindingVersion
	return obj
}



// MarshalYAML is a custom marshaller that converts ServerBinding to YAML
func (t ServerBinding) MarshalYAML() (interface{}, error) {
    bytes, err := json.Marshal(t)
    if err != nil {
        return nil, err
    }
    var out interface{}
    err = yaml.Unmarshal(bytes, &out) 
    return out, err
}

// UnmarshalYAML is a custom unmarshaler that converts YAML to ServerBinding
func (t *ServerBinding) UnmarshalYAML(unmarshal func(interface{}) error) error {
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

// MarshalJSON is a custom marshaller that converts ServerBinding to JSON
func (t ServerBinding) MarshalJSON() ([]byte, error) {
	type Alias ServerBinding
	return json.Marshal(struct{ Alias }{Alias(t)})
}

// UnmarshalJSON is a custom unmarshaler that converts JSON to ServerBinding
func (t *ServerBinding) UnmarshalJSON(data []byte) error {
	type Alias ServerBinding
	aux := struct{ *Alias }{Alias: (*Alias)(t)}
	return json.Unmarshal(data, &aux)
}


