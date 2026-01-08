# asyncapi-go

Go library for parsing and working with [AsyncAPI](https://www.asyncapi.com/) specifications. Supports both AsyncAPI 2.x and 3.x versions.

## Supported Versions

| Version | Status |
|---------|--------|
| 2.0.0 - 2.6.0 | Supported |
| 3.0.0 | Supported |

## Supported Bindings

All official AsyncAPI bindings with defined properties are supported:

| Binding | Version | Description |
|---------|---------|-------------|
| amqp | 0.3.0 | RabbitMQ / AMQP 0-9-1 |
| anypointmq | 0.1.0 | MuleSoft Anypoint MQ |
| googlepubsub | 0.2.0 | Google Cloud Pub/Sub |
| http | 0.3.0 | HTTP/REST |
| ibmmq | 0.1.0 | IBM MQ |
| jms | 0.0.1 | Java Message Service |
| kafka | 0.5.0 | Apache Kafka |
| mqtt | 0.2.0 | MQTT 3.x/5.x |
| nats | 0.1.0 | NATS |
| pulsar | 0.1.0 | Apache Pulsar |
| sns | 0.3.0 | AWS SNS |
| solace | 0.4.0 | Solace PubSub+ |
| sqs | 0.3.0 | AWS SQS |
| websockets | 0.1.0 | WebSockets |

## Installation

```bash
go get github.com/charlie-haley/asyncapi-go
```

## Usage

### Parsing an AsyncAPI Document

```go
package main

import (
    "fmt"
    "github.com/charlie-haley/asyncapi-go"
)

func main() {
    // Parse from file
    doc, err := asyncapi.ParseFile("asyncapi.yaml")
    if err != nil {
        panic(err)
    }

    fmt.Printf("Title: %s\n", doc.GetTitle())
    fmt.Printf("Version: %s\n", doc.GetVersion())
}
```

### Working with AsyncAPI v3

```go
package main

import (
    "fmt"
    "github.com/charlie-haley/asyncapi-go"
    "github.com/charlie-haley/asyncapi-go/asyncapi3"
    "github.com/charlie-haley/asyncapi-go/bindings/kafka"
)

func main() {
    doc, _ := asyncapi.ParseFile("asyncapi-v3.yaml")

    // Type assert to v3 document
    v3Doc := doc.(*asyncapi3.Document)

    // Access channels
    for name, channel := range v3Doc.Channels {
        fmt.Printf("Channel: %s (address: %s)\n", name, channel.Address)

        // Parse Kafka binding if present
        if channel.HasBinding("kafka") {
            binding, _ := asyncapi.ParseBindings[kafka.ChannelBinding](channel.Bindings, "kafka")
            fmt.Printf("  Kafka topic: %s, partitions: %d\n", binding.Topic, binding.Partitions)
        }
    }

    // Access operations (v3 specific)
    for name, op := range v3Doc.Operations {
        fmt.Printf("Operation: %s (action: %s)\n", name, op.Action)
    }
}
```

### Working with AsyncAPI v2

```go
package main

import (
    "fmt"
    "github.com/charlie-haley/asyncapi-go"
    "github.com/charlie-haley/asyncapi-go/asyncapi2"
    "github.com/charlie-haley/asyncapi-go/bindings/kafka"
)

func main() {
    doc, _ := asyncapi.ParseFile("asyncapi-v2.yaml")

    // Type assert to v2 document
    v2Doc := doc.(*asyncapi2.Document)

    // Access channels
    for name, channel := range v2Doc.Channels {
        fmt.Printf("Channel: %s\n", name)

        // Parse Kafka binding
        if binding, err := asyncapi.ParseBindings[kafka.ChannelBinding](channel.Bindings, "kafka"); err == nil {
            fmt.Printf("  Topic: %s\n", binding.Topic)
        }
    }
}
```

### Parsing Protocol Bindings

Bindings can be parsed from channels, operations, messages, and servers:

```go
import (
    "github.com/charlie-haley/asyncapi-go"
    "github.com/charlie-haley/asyncapi-go/bindings/kafka"
    "github.com/charlie-haley/asyncapi-go/bindings/amqp"
    "github.com/charlie-haley/asyncapi-go/bindings/sqs"
)

// Kafka channel binding
kafkaBinding, _ := asyncapi.ParseBindings[kafka.ChannelBinding](channel.Bindings, "kafka")
fmt.Printf("Topic: %s, Partitions: %d\n", kafkaBinding.Topic, kafkaBinding.Partitions)

// AMQP channel binding
amqpBinding, _ := asyncapi.ParseBindings[amqp.ChannelBinding](channel.Bindings, "amqp")
fmt.Printf("Exchange: %s, Queue: %s\n", amqpBinding.Exchange.Name, amqpBinding.Queue.Name)

// SQS channel binding
sqsBinding, _ := asyncapi.ParseBindings[sqs.ChannelBinding](channel.Bindings, "sqs")
fmt.Printf("Queue: %s, FIFO: %v\n", sqsBinding.Queue.Name, sqsBinding.Queue.FifoQueue)
```

### Custom Bindings

You can parse custom bindings by defining your own struct:

```go
type CustomBinding struct {
    CustomField string `json:"customField"`
    Options     []string `json:"options"`
}

binding, _ := asyncapi.ParseBindings[CustomBinding](channel.Bindings, "custom")
```

## Reference Resolution

The library automatically resolves `$ref` references within documents, including:
- Local references (`#/components/messages/UserMessage`)
- File references (`./common/schemas.yaml`)
- Remote references (`https://example.com/schemas.json`)

## License

MIT
