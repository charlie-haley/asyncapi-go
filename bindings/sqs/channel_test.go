package sqs

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestChannelBinding_BuildObject(t *testing.T) {
	maxReceiveCount := 3
	cb := NewChannelBinding().
		WithQueue(NewQueue().
			WithName("MyQueue").
			WithFifoQueue(true).
			WithDeduplicationScope("messageGroup").
			WithFifoThroughputLimit("perMessageGroupId").
			WithDeliveryDelay(60).
			WithVisibilityTimeout(30).
			WithReceiveMessageWaitTime(20).
			WithMessageRetentionPeriod(345600).
			WithRedrivePolicy(&RedrivePolicy{
				DeadLetterQueue: Identifier{
					ARN:  "arn:aws:sqs:us-east-1:123456789012:MyDeadLetterQueue",
					Name: "MyDeadLetterQueue",
				},
				MaxReceiveCount: &maxReceiveCount,
			}).
			WithPolicy(&Policy{
				Statements: []Statement{
					{
						Effect:    "Allow",
						Principal: "*",
						Action:    "sqs:SendMessage",
						Resource:  "arn:aws:sqs:us-east-1:123456789012:MyQueue",
					},
				},
			}).
			WithTags(map[string]string{
				"environment": "production",
				"team":        "platform",
			}),
		).
		WithDeadLetterQueue(NewQueue().
			WithName("MyDeadLetterQueue").
			WithFifoQueue(true))

	assert.NotNil(t, cb.Queue)
	assert.Equal(t, "MyQueue", cb.Queue.Name)
	assert.True(t, cb.Queue.FifoQueue)
	assert.Equal(t, "messageGroup", cb.Queue.DeduplicationScope)
	assert.Equal(t, "perMessageGroupId", cb.Queue.FifoThroughputLimit)
	assert.Equal(t, 60, cb.Queue.DeliveryDelay)
	assert.Equal(t, 30, cb.Queue.VisibilityTimeout)
	assert.Equal(t, 20, cb.Queue.ReceiveMessageWaitTime)
	assert.Equal(t, 345600, cb.Queue.MessageRetentionPeriod)
	assert.NotNil(t, cb.Queue.RedrivePolicy)
	assert.Equal(t, "arn:aws:sqs:us-east-1:123456789012:MyDeadLetterQueue", cb.Queue.RedrivePolicy.DeadLetterQueue.ARN)
	assert.Equal(t, "MyDeadLetterQueue", cb.Queue.RedrivePolicy.DeadLetterQueue.Name)
	assert.Equal(t, 3, *cb.Queue.RedrivePolicy.MaxReceiveCount)
	assert.NotNil(t, cb.Queue.Policy)
	assert.Len(t, cb.Queue.Policy.Statements, 1)
	assert.Equal(t, "Allow", cb.Queue.Policy.Statements[0].Effect)
	assert.Equal(t, "*", cb.Queue.Policy.Statements[0].Principal)
	assert.Equal(t, "sqs:SendMessage", cb.Queue.Policy.Statements[0].Action)
	assert.Equal(t, "arn:aws:sqs:us-east-1:123456789012:MyQueue", cb.Queue.Policy.Statements[0].Resource)
	assert.Equal(t, "production", cb.Queue.Tags["environment"])
	assert.Equal(t, "platform", cb.Queue.Tags["team"])
	assert.NotNil(t, cb.DeadLetterQueue)
	assert.Equal(t, "MyDeadLetterQueue", cb.DeadLetterQueue.Name)
	assert.True(t, cb.DeadLetterQueue.FifoQueue)
}

func TestChannelBinding_MarshalYAML(t *testing.T) {
	maxReceiveCount := 3
	cb := NewChannelBinding().
		WithQueue(NewQueue().
			WithName("MyQueue").
			WithFifoQueue(true).
			WithRedrivePolicy(&RedrivePolicy{
				DeadLetterQueue: Identifier{
					Name: "MyDeadLetterQueue",
				},
				MaxReceiveCount: &maxReceiveCount,
			}))

	expectedYAML := `bindingVersion: latest
queue:
  fifoQueue: true
  name: MyQueue
  redrivePolicy:
    deadLetterQueue:
      name: MyDeadLetterQueue
    maxReceiveCount: 3
`
	marshaledYAML, err := yaml.Marshal(cb)
	assert.NoError(t, err)
	assert.Equal(t, expectedYAML, string(marshaledYAML))
}

func TestChannelBinding_UnmarshalYAML(t *testing.T) {
	yamlString := `
queue:
  name: MyQueue
  fifoQueue: true
  deduplicationScope: messageGroup
  fifoThroughputLimit: perMessageGroupId
  deliveryDelay: 60
  visibilityTimeout: 30
  receiveMessageWaitTime: 20
  messageRetentionPeriod: 345600
  redrivePolicy:
    deadLetterQueue:
      name: MyDeadLetterQueue
    maxReceiveCount: 3
  tags:
    environment: production
deadLetterQueue:
  name: MyDeadLetterQueue
  fifoQueue: true
bindingVersion: "0.5.0"
`
	var cb ChannelBinding
	err := yaml.Unmarshal([]byte(yamlString), &cb)
	assert.NoError(t, err)

	assert.NotNil(t, cb.Queue)
	assert.Equal(t, "MyQueue", cb.Queue.Name)
	assert.True(t, cb.Queue.FifoQueue)
	assert.Equal(t, "messageGroup", cb.Queue.DeduplicationScope)
	assert.Equal(t, "perMessageGroupId", cb.Queue.FifoThroughputLimit)
	assert.Equal(t, 60, cb.Queue.DeliveryDelay)
	assert.Equal(t, 30, cb.Queue.VisibilityTimeout)
	assert.Equal(t, 20, cb.Queue.ReceiveMessageWaitTime)
	assert.Equal(t, 345600, cb.Queue.MessageRetentionPeriod)
	assert.NotNil(t, cb.Queue.RedrivePolicy)
	assert.Equal(t, "MyDeadLetterQueue", cb.Queue.RedrivePolicy.DeadLetterQueue.Name)
	assert.Equal(t, 3, *cb.Queue.RedrivePolicy.MaxReceiveCount)
	assert.Equal(t, "production", cb.Queue.Tags["environment"])
	assert.NotNil(t, cb.DeadLetterQueue)
	assert.Equal(t, "MyDeadLetterQueue", cb.DeadLetterQueue.Name)
	assert.True(t, cb.DeadLetterQueue.FifoQueue)
}

func TestChannelBinding_MarshalJSON(t *testing.T) {
	maxReceiveCount := 3
	cb := NewChannelBinding().
		WithQueue(NewQueue().
			WithName("MyQueue").
			WithFifoQueue(true).
			WithRedrivePolicy(&RedrivePolicy{
				DeadLetterQueue: Identifier{
					Name: "MyDeadLetterQueue",
				},
				MaxReceiveCount: &maxReceiveCount,
			}))

	expectedJSON := `{"queue":{"name":"MyQueue","fifoQueue":true,"redrivePolicy":{"deadLetterQueue":{"name":"MyDeadLetterQueue"},"maxReceiveCount":3}},"bindingVersion":"latest"}`

	marshaledJSON, err := json.Marshal(cb)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(marshaledJSON))
}

func TestChannelBinding_UnmarshalJSON(t *testing.T) {
	jsonString := `{
		"queue": {
			"name": "MyQueue",
			"fifoQueue": true,
			"deduplicationScope": "messageGroup",
			"fifoThroughputLimit": "perMessageGroupId",
			"deliveryDelay": 60,
			"visibilityTimeout": 30,
			"receiveMessageWaitTime": 20,
			"messageRetentionPeriod": 345600,
			"redrivePolicy": {
				"deadLetterQueue": {
					"name": "MyDeadLetterQueue"
				},
				"maxReceiveCount": 3
			},
			"policy": {
				"statements": [
					{
						"effect": "Allow",
						"principal": "*",
						"action": "sqs:SendMessage"
					}
				]
			},
			"tags": {
				"environment": "production"
			}
		}
	}`

	var cb ChannelBinding
	err := json.Unmarshal([]byte(jsonString), &cb)
	assert.NoError(t, err)

	assert.NotNil(t, cb.Queue)
	assert.Equal(t, "MyQueue", cb.Queue.Name)
	assert.True(t, cb.Queue.FifoQueue)
	assert.Equal(t, "messageGroup", cb.Queue.DeduplicationScope)
	assert.Equal(t, "perMessageGroupId", cb.Queue.FifoThroughputLimit)
	assert.Equal(t, 60, cb.Queue.DeliveryDelay)
	assert.Equal(t, 30, cb.Queue.VisibilityTimeout)
	assert.Equal(t, 20, cb.Queue.ReceiveMessageWaitTime)
	assert.Equal(t, 345600, cb.Queue.MessageRetentionPeriod)
	assert.NotNil(t, cb.Queue.RedrivePolicy)
	assert.Equal(t, "MyDeadLetterQueue", cb.Queue.RedrivePolicy.DeadLetterQueue.Name)
	assert.Equal(t, 3, *cb.Queue.RedrivePolicy.MaxReceiveCount)
	assert.NotNil(t, cb.Queue.Policy)
	assert.Len(t, cb.Queue.Policy.Statements, 1)
	assert.Equal(t, "Allow", cb.Queue.Policy.Statements[0].Effect)
	assert.Equal(t, "*", cb.Queue.Policy.Statements[0].Principal)
	assert.Equal(t, "sqs:SendMessage", cb.Queue.Policy.Statements[0].Action)
	assert.Equal(t, "production", cb.Queue.Tags["environment"])
}