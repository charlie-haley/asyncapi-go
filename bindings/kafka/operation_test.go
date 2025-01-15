package kafka

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestOperationBinding_BuildObject(t *testing.T) {
	ob := NewOperationBinding().
		WithGroupID("test-group").
		WithClientID("test-client")

	assert.Equal(t, "test-group", ob.GroupID)
	assert.Equal(t, "test-client", ob.ClientID)
}

func TestOperationBinding_MarshalYAML(t *testing.T) {
	ob := NewOperationBinding().
		WithGroupID("test-group").
		WithClientID("test-client")

	expectedYAML := `clientId: test-client
groupId: test-group
`
	marshaledYAML, err := yaml.Marshal(ob)
	assert.NoError(t, err)
	assert.Equal(t, expectedYAML, string(marshaledYAML))
}

func TestOperationBinding_UnmarshalYAML(t *testing.T) {
	yamlString := `
groupId: test-group
clientId: test-client
bindingVersion: "0.5.0"
`
	var ob OperationBinding
	err := yaml.Unmarshal([]byte(yamlString), &ob)
	assert.NoError(t, err)

	assert.Equal(t, "test-group", ob.GroupID)
	assert.Equal(t, "test-client", ob.ClientID)
	assert.Equal(t, "0.5.0", ob.BindingVersion)
}

func TestOperationBinding_MarshalJSON(t *testing.T) {
	ob := NewOperationBinding().
		WithGroupID("test-group").
		WithClientID("test-client")

	expectedJSON := `{"groupId":"test-group","clientId":"test-client"}`

	marshaledJSON, err := json.Marshal(ob)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(marshaledJSON))
}

func TestOperationBinding_UnmarshalJSON(t *testing.T) {
	jsonString := `{
		"groupId": "test-group",
		"clientId": "test-client",
		"bindingVersion": "0.5.0"
	}`

	var ob OperationBinding
	err := json.Unmarshal([]byte(jsonString), &ob)
	assert.NoError(t, err)

	assert.Equal(t, "test-group", ob.GroupID)
	assert.Equal(t, "test-client", ob.ClientID)
	assert.Equal(t, "0.5.0", ob.BindingVersion)
}

func TestOperationBinding_NonStringFields(t *testing.T) {
	jsonString := `{
		"groupId": 123,
		"clientId": ["client1", "client2"],
		"bindingVersion": "0.5.0"
	}`

	var ob OperationBinding
	err := json.Unmarshal([]byte(jsonString), &ob)
	assert.NoError(t, err)
	assert.Equal(t, float64(123), ob.GroupID)
	assert.Equal(t, []interface{}{"client1", "client2"}, ob.ClientID)
}
