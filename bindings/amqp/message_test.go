package amqp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestMessageBinding_BuildObject(t *testing.T) {
	mb := NewMessageBinding().
		WithContentEncoding("gzip").
		WithMessageType("user.signup")

	assert.Equal(t, "gzip", mb.ContentEncoding)
	assert.Equal(t, "user.signup", mb.MessageType)
}

func TestMessageBinding_MarshalYAML(t *testing.T) {
	mb := NewMessageBinding().
		WithContentEncoding("gzip").
		WithMessageType("user.signup")

	expectedYAML := `contentEncoding: gzip
messageType: user.signup
`
	marshaledYAML, err := yaml.Marshal(mb)
	assert.NoError(t, err)
	assert.Equal(t, expectedYAML, string(marshaledYAML))
}

func TestMessageBinding_UnmarshalYAML(t *testing.T) {
	yamlString := `
contentEncoding: application/json
messageType: user.login
`

	var mb MessageBinding
	err := yaml.Unmarshal([]byte(yamlString), &mb)
	assert.NoError(t, err)

	assert.Equal(t, "application/json", mb.ContentEncoding)
	assert.Equal(t, "user.login", mb.MessageType)
}
