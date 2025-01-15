package amqp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestOperationBindingConstants(t *testing.T) {
	assert.Equal(t, 1, OperationDeliveryModeTransient)
	assert.Equal(t, 2, OperationDeliveryModePersistent)
}

func TestOperationBinding_BuildObject(t *testing.T) {
	ob := NewOperationBinding().
		WithExpiration(1000).
		WithUserID("guest").
		WithCC([]string{"user.logs"}).
		WithPriority(1).
		WithDeliveryMode(OperationDeliveryModePersistent).
		WithMandatory(true).
		WithBCC([]string{"external.audit"}).
		WithTimestamp(true).
		WithAck(false)

	assert.Equal(t, 1000, ob.Expiration)
	assert.Equal(t, "guest", ob.UserID)
	assert.Equal(t, []string{"user.logs"}, ob.CC)
	assert.Equal(t, 1, ob.Priority)
	assert.Equal(t, OperationDeliveryModePersistent, ob.DeliveryMode)
	assert.True(t, ob.Mandatory)
	assert.Equal(t, []string{"external.audit"}, ob.BCC)
	assert.True(t, ob.Timestamp)
	assert.False(t, ob.Ack)
}

func TestOperationBinding_MarshalYAML(t *testing.T) {
	ob := NewOperationBinding().
		WithExpiration(1000).
		WithUserID("guest").
		WithCC([]string{"user.logs"}).
		WithPriority(1).
		WithDeliveryMode(OperationDeliveryModePersistent).
		WithMandatory(true).
		WithBCC([]string{"external.audit"}).
		WithTimestamp(true).
		WithAck(false)

	expectedYAML := `bcc:
- external.audit
cc:
- user.logs
deliveryMode: 2
expiration: 1000
mandatory: true
priority: 1
timestamp: true
userId: guest
`

	marshaledYAML, err := yaml.Marshal(ob)
	assert.NoError(t, err)
	assert.Equal(t, expectedYAML, string(marshaledYAML))
}

func TestOperationBinding_UnmarshalYAML(t *testing.T) {
	yamlString := `
expiration: 2000
userId: admin
cc:
  - user.notifications
priority: 2
deliveryMode: 1
mandatory: false
bcc:
  - internal.audit
timestamp: false
ack: true
`

	var ob OperationBinding
	err := yaml.Unmarshal([]byte(yamlString), &ob)
	assert.NoError(t, err)

	assert.Equal(t, 2000, ob.Expiration)
	assert.Equal(t, "admin", ob.UserID)
	assert.Equal(t, []string{"user.notifications"}, ob.CC)
	assert.Equal(t, 2, ob.Priority)
	assert.Equal(t, OperationDeliveryModeTransient, ob.DeliveryMode)
	assert.False(t, ob.Mandatory)
	assert.Equal(t, []string{"internal.audit"}, ob.BCC)
	assert.False(t, ob.Timestamp)
	assert.True(t, ob.Ack)
}
