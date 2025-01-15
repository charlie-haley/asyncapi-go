package sns

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestOperationBinding_YAML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected OperationBinding
	}{
		{
			name: "full configuration",
			input: `
topic:
  arn: arn:aws:sns:us-west-2:123456789012:MyTopic
consumers:
  - protocol: sqs
    endpoint:
      arn: arn:aws:sqs:us-west-2:123456789012:MyQueue
    filterPolicy:
      event:
        - order_created
        - order_updated
    filterPolicyScope: MessageBody
    rawMessageDelivery: true
    redrivePolicy:
      deadLetterQueue:
        arn: arn:aws:sqs:us-west-2:123456789012:MyDLQ
      maxReceiveCount: 5
    deliveryPolicy:
      minDelayTarget: 1
      maxDelayTarget: 60
      numRetries: 50
      numNoDelayRetries: 3
      backoffFunction: exponential
    displayName: Orders Queue
deliveryPolicy:
  minDelayTarget: 2
  maxDelayTarget: 120
  numRetries: 100
  maxReceivesPerSecond: 10
bindingVersion: "1.0.0"
`,
			expected: OperationBinding{
				Topic: &Identifier{
					ARN: "arn:aws:sns:us-west-2:123456789012:MyTopic",
				},
				Consumers: []Consumer{
					{
						Protocol: "sqs",
						Endpoint: Identifier{
							ARN: "arn:aws:sqs:us-west-2:123456789012:MyQueue",
						},
						FilterPolicy: map[string]interface{}{
							"event": []interface{}{"order_created", "order_updated"},
						},
						FilterPolicyScope:  "MessageBody",
						RawMessageDelivery: true,
						RedrivePolicy: &RedrivePolicy{
							DeadLetterQueue: Identifier{
								ARN: "arn:aws:sqs:us-west-2:123456789012:MyDLQ",
							},
							MaxReceiveCount: intPtr(5),
						},
						DeliveryPolicy: &DeliveryPolicy{
							MinDelayTarget:    intPtr(1),
							MaxDelayTarget:    intPtr(60),
							NumRetries:        intPtr(50),
							NumNoDelayRetries: intPtr(3),
							BackoffFunction:   "exponential",
						},
						DisplayName: "Orders Queue",
					},
				},
				DeliveryPolicy: &DeliveryPolicy{
					MinDelayTarget:       intPtr(2),
					MaxDelayTarget:       intPtr(120),
					NumRetries:           intPtr(100),
					MaxReceivesPerSecond: intPtr(10),
				},
				BindingVersion: "1.0.0",
			},
		},
		{
			name: "multiple protocols",
			input: `consumers:
  - protocol: http
    endpoint:
      url: https://example.com/webhook
    rawMessageDelivery: true
  - protocol: email
    endpoint:
      email: test@example.com
    rawMessageDelivery: false
  - protocol: sms
    endpoint:
      phone: "+1234567890"
    rawMessageDelivery: true
    displayName: SMS Alerts`,
			expected: OperationBinding{
				Consumers: []Consumer{
					{
						Protocol: "http",
						Endpoint: Identifier{
							URL: "https://example.com/webhook",
						},
						RawMessageDelivery: true,
					},
					{
						Protocol: "email",
						Endpoint: Identifier{
							Email: "test@example.com",
						},
						RawMessageDelivery: false,
					},
					{
						Protocol: "sms",
						Endpoint: Identifier{
							Phone: "+1234567890",
						},
						RawMessageDelivery: true,
						DisplayName:       "SMS Alerts",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got OperationBinding
			err := yaml.Unmarshal([]byte(tt.input), &got)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got)

			// Test marshaling
			marshaled, err := yaml.Marshal(got)
			assert.NoError(t, err)

			var unmarshaled OperationBinding
			err = yaml.Unmarshal(marshaled, &unmarshaled)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, unmarshaled)
		})
	}
}

func TestOperationBinding_JSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected OperationBinding
	}{
		{
			name: "full configuration",
			input: `{
				"topic": {
					"arn": "arn:aws:sns:us-west-2:123456789012:MyTopic"
				},
				"consumers": [
					{
						"protocol": "sqs",
						"endpoint": {
							"arn": "arn:aws:sqs:us-west-2:123456789012:MyQueue"
						},
						"filterPolicy": {
							"event": ["order_created", "order_updated"]
						},
						"filterPolicyScope": "MessageBody",
						"rawMessageDelivery": true,
						"redrivePolicy": {
							"deadLetterQueue": {
								"arn": "arn:aws:sqs:us-west-2:123456789012:MyDLQ"
							},
							"maxReceiveCount": 5
						},
						"deliveryPolicy": {
							"minDelayTarget": 1,
							"maxDelayTarget": 60,
							"numRetries": 50,
							"numNoDelayRetries": 3,
							"backoffFunction": "exponential"
						},
						"displayName": "Orders Queue"
					}
				],
				"deliveryPolicy": {
					"minDelayTarget": 2,
					"maxDelayTarget": 120,
					"numRetries": 100,
					"maxReceivesPerSecond": 10
				},
				"bindingVersion": "1.0.0"
			}`,
			expected: OperationBinding{
				Topic: &Identifier{
					ARN: "arn:aws:sns:us-west-2:123456789012:MyTopic",
				},
				Consumers: []Consumer{
					{
						Protocol: "sqs",
						Endpoint: Identifier{
							ARN: "arn:aws:sqs:us-west-2:123456789012:MyQueue",
						},
						FilterPolicy: map[string]interface{}{
							"event": []interface{}{"order_created", "order_updated"},
						},
						FilterPolicyScope:  "MessageBody",
						RawMessageDelivery: true,
						RedrivePolicy: &RedrivePolicy{
							DeadLetterQueue: Identifier{
								ARN: "arn:aws:sqs:us-west-2:123456789012:MyDLQ",
							},
							MaxReceiveCount: intPtr(5),
						},
						DeliveryPolicy: &DeliveryPolicy{
							MinDelayTarget:    intPtr(1),
							MaxDelayTarget:    intPtr(60),
							NumRetries:        intPtr(50),
							NumNoDelayRetries: intPtr(3),
							BackoffFunction:   "exponential",
						},
						DisplayName: "Orders Queue",
					},
				},
				DeliveryPolicy: &DeliveryPolicy{
					MinDelayTarget:       intPtr(2),
					MaxDelayTarget:       intPtr(120),
					NumRetries:           intPtr(100),
					MaxReceivesPerSecond: intPtr(10),
				},
				BindingVersion: "1.0.0",
			},
		},
		{
			name: "multiple protocols",
			input: `{
				"consumers": [
					{
						"protocol": "http",
						"endpoint": {
							"url": "https://example.com/webhook"
						},
						"rawMessageDelivery": true
					},
					{
						"protocol": "email",
						"endpoint": {
							"email": "test@example.com"
						},
						"rawMessageDelivery": false
					},
					{
						"protocol": "sms",
						"endpoint": {
							"phone": "+1234567890"
						},
						"rawMessageDelivery": true,
						"displayName": "SMS Alerts"
					}
				]
			}`,
			expected: OperationBinding{
				Consumers: []Consumer{
					{
						Protocol: "http",
						Endpoint: Identifier{
							URL: "https://example.com/webhook",
						},
						RawMessageDelivery: true,
					},
					{
						Protocol: "email",
						Endpoint: Identifier{
							Email: "test@example.com",
						},
						RawMessageDelivery: false,
					},
					{
						Protocol: "sms",
						Endpoint: Identifier{
							Phone: "+1234567890",
						},
						RawMessageDelivery: true,
						DisplayName:       "SMS Alerts",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got OperationBinding
			err := json.Unmarshal([]byte(tt.input), &got)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got)

			// Test marshaling
			marshaled, err := json.Marshal(got)
			assert.NoError(t, err)

			var unmarshaled OperationBinding
			err = json.Unmarshal(marshaled, &unmarshaled)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, unmarshaled)
		})
	}
}

// Helper function for creating int pointers
func intPtr(i int) *int {
	return &i
}