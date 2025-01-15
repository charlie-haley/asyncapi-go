package sqs

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestOperationBinding_BuildObject(t *testing.T) {
	ob := NewOperationBinding().
		WithQueues([]Queue{
			*NewQueue().
				WithName("Queue1").
				WithFifoQueue(true),
			*NewQueue().
				WithName("Queue2").
				WithFifoQueue(false),
		})

	assert.Len(t, ob.Queues, 2)
	assert.Equal(t, "Queue1", ob.Queues[0].Name)
	assert.True(t, ob.Queues[0].FifoQueue)
	assert.Equal(t, "Queue2", ob.Queues[1].Name)
	assert.False(t, ob.Queues[1].FifoQueue)
}

func TestOperationBinding_MarshalYAML(t *testing.T) {
	ob := NewOperationBinding().
		WithQueues([]Queue{
			*NewQueue().
				WithName("Queue1").
				WithFifoQueue(true),
		})

	expectedYAML := `queues:
- fifoQueue: true
  name: Queue1
`
	marshaledYAML, err := yaml.Marshal(ob)
	assert.NoError(t, err)
	assert.Equal(t, expectedYAML, string(marshaledYAML))
}

func TestOperationBinding_UnmarshalYAML(t *testing.T) {
	yamlString := `
queues:
- name: Queue1
  fifoQueue: true
- name: Queue2
  fifoQueue: false
`
	var ob OperationBinding
	err := yaml.Unmarshal([]byte(yamlString), &ob)
	assert.NoError(t, err)

	assert.Len(t, ob.Queues, 2)
	assert.Equal(t, "Queue1", ob.Queues[0].Name)
	assert.True(t, ob.Queues[0].FifoQueue)
	assert.Equal(t, "Queue2", ob.Queues[1].Name)
	assert.False(t, ob.Queues[1].FifoQueue)
}

func TestOperationBinding_MarshalJSON(t *testing.T) {
	ob := NewOperationBinding().
		WithQueues([]Queue{
			*NewQueue().
				WithName("Queue1").
				WithFifoQueue(true),
		})

	expectedJSON := `{"queues":[{"name":"Queue1","fifoQueue":true}]}`

	marshaledJSON, err := json.Marshal(ob)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(marshaledJSON))
}

func TestOperationBinding_UnmarshalJSON(t *testing.T) {
	jsonString := `{
		"queues": [
			{
				"name": "Queue1",
				"fifoQueue": true
			},
			{
				"name": "Queue2",
				"fifoQueue": false
			}
		]
	}`

	var ob OperationBinding
	err := json.Unmarshal([]byte(jsonString), &ob)
	assert.NoError(t, err)

	assert.Len(t, ob.Queues, 2)
	assert.Equal(t, "Queue1", ob.Queues[0].Name)
	assert.True(t, ob.Queues[0].FifoQueue)
	assert.Equal(t, "Queue2", ob.Queues[1].Name)
	assert.False(t, ob.Queues[1].FifoQueue)
}