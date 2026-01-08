package mqtt

// MessageBinding represents the MQTT Message Binding object.
//
// This object contains information about the message representation in MQTT.
// +binding
type MessageBinding struct {
	// PayloadFormatIndicator indicates the format of the payload.
	// 0 for unspecified bytes, 1 for UTF-8 encoded character data.
	PayloadFormatIndicator int `json:"payloadFormatIndicator,omitempty"`

	// CorrelationData is used for request-response correlation.
	// Can be a Schema Object or Reference Object.
	CorrelationData interface{} `json:"correlationData,omitempty"`

	// ContentType describes the content type of the message payload.
	ContentType string `json:"contentType,omitempty"`

	// ResponseTopic is the topic name for a response message.
	// Can be a string, Schema Object, or Reference Object.
	ResponseTopic interface{} `json:"responseTopic,omitempty"`

	// BindingVersion is the version of this binding.
	BindingVersion string `json:"bindingVersion,omitempty"`
}
