package amqp

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/yaml"
)

func TestChannelBindingConstants(t *testing.T) {
	assert.Equal(t, "routingKey", ChannelIsRoutingKey)
	assert.Equal(t, "queue", ChannelIsQueue)
}

func TestExchangeConstants(t *testing.T) {
	assert.Equal(t, "topic", ExchangeTypeTopic)
	assert.Equal(t, "direct", ExchangeTypeDirect)
	assert.Equal(t, "fanout", ExchangeTypeFanout)
	assert.Equal(t, "default", ExchangeTypeDefault)
	assert.Equal(t, "headers", ExchangeTypeHeaders)
}

func TestChannelBinding_BuildObject(t *testing.T) {
	cb := NewChannelBinding().
		WithIs(ChannelIsQueue).
		WithExchange(NewExchange().
			WithName("testExchange").
			WithType(ExchangeTypeDirect).
			WithDurable(true).
			WithAutoDelete(false).
			WithVHost("/"),
		).
		WithQueue(NewQueue().
			WithName("testQueue").
			WithDurable(true).
			WithExclusive(true).
			WithAutoDelete(false).
			WithVHost("/"),
		)

	assert.Equal(t, ChannelIsQueue, cb.Is)
	assert.NotNil(t, cb.Exchange)
	assert.Equal(t, "testExchange", cb.Exchange.Name)
	assert.Equal(t, ExchangeTypeDirect, cb.Exchange.Type)
	assert.True(t, cb.Exchange.Durable)
	assert.False(t, cb.Exchange.AutoDelete)
	assert.Equal(t, "/", cb.Exchange.VHost)
	assert.NotNil(t, cb.Queue)
	assert.Equal(t, "testQueue", cb.Queue.Name)
	assert.True(t, cb.Queue.Durable)
	assert.True(t, cb.Queue.Exclusive)
	assert.False(t, cb.Queue.AutoDelete)
	assert.Equal(t, "/", cb.Queue.VHost)
}

func TestChannelBinding_MarshalYAML(t *testing.T) {
	cb := NewChannelBinding().
		WithIs(ChannelIsQueue).
		WithExchange(NewExchange().
			WithName("testExchange").
			WithType(ExchangeTypeDirect).
			WithDurable(true).
			WithAutoDelete(false).
			WithVHost("/"),
		).
		WithQueue(NewQueue().
			WithName("testQueue").
			WithDurable(true).
			WithExclusive(true).
			WithAutoDelete(false).
			WithVHost("/"),
		)

	expectedYAML := `exchange:
  autoDelete: false
  durable: true
  name: testExchange
  type: direct
  vhost: /
is: queue
queue:
  autoDelete: false
  durable: true
  exclusive: true
  name: testQueue
  vhost: /
`

	marshaledYAML, err := yaml.Marshal(cb)
	assert.NoError(t, err)
	assert.Equal(t, expectedYAML, string(marshaledYAML))
}

func TestChannelBinding_UnmarshalYAML(t *testing.T) {
	yamlString := `
is: routingKey
exchange:
  name: testExchange
  type: headers
  durable: false
  autoDelete: true
  vhost: /test
queue:
  name: testQueue
  durable: false
  exclusive: false
  autoDelete: true
  vhost: /test
`

	var cb ChannelBinding
	err := yaml.Unmarshal([]byte(yamlString), &cb)
	assert.NoError(t, err)

	assert.Equal(t, ChannelIsRoutingKey, cb.Is)
	assert.NotNil(t, cb.Exchange)
	assert.Equal(t, "testExchange", cb.Exchange.Name)
	assert.Equal(t, ExchangeTypeHeaders, cb.Exchange.Type)
	assert.False(t, cb.Exchange.Durable)
	assert.True(t, cb.Exchange.AutoDelete)
	assert.Equal(t, "/test", cb.Exchange.VHost)
	assert.NotNil(t, cb.Queue)
	assert.Equal(t, "testQueue", cb.Queue.Name)
	assert.False(t, cb.Queue.Durable)
	assert.False(t, cb.Queue.Exclusive)
	assert.True(t, cb.Queue.AutoDelete)
	assert.Equal(t, "/test", cb.Queue.VHost)
}

func TestChannelBinding_MarshalJSON(t *testing.T) {
	cb := NewChannelBinding().
		WithIs(ChannelIsQueue).
		WithExchange(NewExchange().
			WithName("testExchange").
			WithType(ExchangeTypeDirect).
			WithDurable(true).
			WithAutoDelete(false).
			WithVHost("/"),
		).
		WithQueue(NewQueue().
			WithName("testQueue").
			WithDurable(true).
			WithExclusive(true).
			WithAutoDelete(false).
			WithVHost("/"),
		)

	expectedJSON := `{"is":"queue","exchange":{"name":"testExchange","type":"direct","durable":true,"autoDelete":false,"vhost":"/"},"queue":{"name":"testQueue","durable":true,"exclusive":true,"autoDelete":false,"vhost":"/"}}`

	marshaledJSON, err := json.Marshal(cb)
	assert.NoError(t, err)
	assert.Equal(t, expectedJSON, string(marshaledJSON))
}

func TestChannelBinding_UnmarshalJSON(t *testing.T) {
	jsonString := `{"is":"routingKey","exchange":{"name":"testExchange","type":"headers","durable":false,"autoDelete":true,"vhost":"/test"},"queue":{"name":"testQueue","durable":false,"exclusive":false,"autoDelete":true,"vhost":"/test"}}`

	var cb ChannelBinding
	err := json.Unmarshal([]byte(jsonString), &cb)
	assert.NoError(t, err)

	assert.Equal(t, ChannelIsRoutingKey, cb.Is)
	assert.NotNil(t, cb.Exchange)
	assert.Equal(t, "testExchange", cb.Exchange.Name)
	assert.Equal(t, ExchangeTypeHeaders, cb.Exchange.Type)
	assert.False(t, cb.Exchange.Durable)
	assert.True(t, cb.Exchange.AutoDelete)
	assert.Equal(t, "/test", cb.Exchange.VHost)
	assert.NotNil(t, cb.Queue)
	assert.Equal(t, "testQueue", cb.Queue.Name)
	assert.False(t, cb.Queue.Durable)
	assert.False(t, cb.Queue.Exclusive)
	assert.True(t, cb.Queue.AutoDelete)
	assert.Equal(t, "/test", cb.Queue.VHost)
}

func TestExchange_WithDefaults(t *testing.T) {
	e := NewExchange()
	assert.Equal(t, "/", e.VHost)
}

func TestQueue_WithDefaults(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, "/", q.VHost)
}
