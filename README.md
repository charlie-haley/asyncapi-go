# asyncapi-go

**NOTICE: This library is currently under active development and is not yet recommended for production use. Functionality may change significantly, and stability is not guaranteed.**

Go library for parsing and working with [AsyncAPI](https://www.asyncapi.com/) specifications. It currently supports AsyncAPI version 2.x, allowing you to load, validate, and access data within your AsyncAPI documents.

## ✅ Supported Versions

Currently, this library supports the following AsyncAPI versions:

- **2.0.0**
- **2.1.0**
- **2.2.0**
- **2.3.0**
- **2.4.0**
- **2.5.0**
- **2.6.0**

## 🔗 Supported Bindings

This library is under active development, and support for all bindings is not yet complete. Currently, the following bindings are supported:

- amqp
- kafka
- sns
- sqs

## 🚀 Usage

### 📑 Parsing an AsyncAPI Document

Here's how to parse a basic AsyncAPI document:

```go
package main

import (
    "fmt"

    "github.com/charlie-haley/asyncapi-go"
)

func main() {
	asyncapiDocument := ]byte(`
asyncapi: '2.6.0'
info:
  title: My API
  version: '1.0.0'
channels:
  example/channel:
    address: 'example/channel'
`)

	doc, _ := asyncapi.Parse(asyncapiDocument)

	fmt.Printf("AsyncAPI Version: %s\n", doc.GetVersion())
}
```

### 🧩 Parsing a Binding

This example demonstrates how to parse a standard Kafka channel binding from a full AsyncAPI document. Let's say we have an AsyncAPI specification that looks like this, with a `kafka` binding in the `channels` section:

```yaml
asyncapi: "2.6.0"
info:
  title: Kafka Example
  version: "1.0.0"
channels:
  user-signup:
    address: "user-signup"
    bindings:
      kafka:
        topic: "my-topic"
        partitions: 20
        replicas: 3
        topicConfiguration:
          cleanup.policy: ["delete", "compact"]
          retention.ms: 604800000
          retention.bytes: 1000000000
          delete.retention.ms: 86400000
          max.message.bytes: 1048588
        bindingVersion: "0.5.0"
```

We can then use this Go code to parse this document from a file named asyncapi.yaml, access the Kafka binding, and print its properties:

```go
package main

import (
	"fmt"
	"os"

    "github.com/charlie-haley/asyncapi-go"
	"github.com/charlie-haley/asyncapi-go/bindings/kafka"
)

func main() {
	filePath := "asyncapi.yaml"
	data, _ := os.ReadFile(filePath)
	doc, _ := main.ParseFromYAML(data)

	channel, _ := doc.Channels["user-signup"]

	kafkaBinding, _ := asyncapi.ParseBindings[kafka.ChannelBinding](channel.Bindings, "kafka")

	fmt.Printf("Kafka Topic: %s\n", kafkaBinding.Topic)
	fmt.Printf("Partitions: %d\n", kafkaBinding.Partitions)
	fmt.Printf("Replicas: %d\n", kafkaBinding.Replicas)
    fmt.Printf("Cleanup Policy: %v\n", kafkaBinding.TopicConfiguration.CleanupPolicy)
    fmt.Printf("Retention (ms): %d\n", kafkaBinding.TopicConfiguration.RetentionMs)
    fmt.Printf("Retention (bytes): %d\n", kafkaBinding.TopicConfiguration.RetentionBytes)
    fmt.Printf("Delete Retention (ms): %d\n", kafkaBinding.TopicConfiguration.DeleteRetentionMs)
    fmt.Printf("Max Message Bytes: %d\n", kafkaBinding.TopicConfiguration.MaxMessageBytes)
	fmt.Printf("Binding Version: %s\n", kafkaBinding.BindingVersion)
}
```

### 🕊️ Parsing a Custom Binding

Let's say you want to extend your AsyncAPI specification with custom information not covered by the standard bindings. AsyncAPI allows you to do this using "bindings." Imagine you've created a specialized binding for a unique protocol, like [IP over Avian Carriers (IPoAC)](https://en.wikipedia.org/wiki/IP_over_Avian_Carriers) and you'd like to parse it into a Go struct.

```yaml
asyncapi: "2.6.0"
info:
  title: IPoAC Example
  version: "1.0.0"
channels:
  pigeon/post:
    address: "pigeon/post"
    publish:
      message:
        payload:
          type: object
          properties:
            messageId:
              type: string
            content:
              type: string
    bindings:
      ipoac:
        carrier: "pigeon"
        defaultRoute: "RFC 1149"
        maxPacketSize: "256 bytes"
        allowedSpecies:
          - "Rock Dove"
          - "Homing Pigeon"
```

If we define a Go struct to represent this IPoAC binding, we can then parse these custom bindings directly from our AsyncAPI document:

```go
package main

import (
	"fmt"
	"log"
	"os"

    "github.com/charlie-haley/asyncapi-go"
)

type IpoacChannelBinding struct {
	Carrier        string   `json:"carrier"`
	DefaultRoute   string   `json:"defaultRoute"`
	MaxPacketSize  string   `json:"maxPacketSize"`
	AllowedSpecies []string `json:"allowedSpecies"`
}

func main() {
	filePath := "asyncapi.yaml"
	data, _ := os.ReadFile(filePath)
	doc, _ := main.ParseFromYAML(data)

	channel, _ := doc.Channels["pigeon/post"]

	ipoacBinding, _ := asyncapi.ParseBindings[IpoacChannelBinding](channel.Bindings, "ipoac")

	fmt.Printf("Carrier: %s\n", ipoacBinding.Carrier)
	fmt.Printf("Default Route: %s\n", ipoacBinding.DefaultRoute)
	fmt.Printf("Max Packet Size: %s\n", ipoacBinding.MaxPacketSize)
	fmt.Printf("Allowed Species: %v\n", ipoacBinding.AllowedSpecies)
}
```

### 🏗️ Creating an AsyncAPI Document

This example shows how to programmatically create an AsyncAPI document with a Kafka channel binding using the fluent API:

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"

    "github.com/charlie-haley/asyncapi-go/asyncapi2"
	"github.com/charlie-haley/asyncapi-go/bindings/kafka"
)

func main() {
	// Create a new AsyncAPI document.
	doc := asyncapi2.NewDocument().
		WithAsyncAPI("2.6.0").
		WithInfo(asyncapi2.NewInfo().
			WithTitle("User Signup API").
			WithVersion("1.0.0")).
		WithServer("kafka-server", asyncapi2.NewServer().
			WithURL("my-kafka-broker:9092").
			WithProtocol("kafka"))

	// Create a Kafka channel binding.
	kafkaBinding := kafka.NewChannelBinding().
		WithTopic("user-signup-topic").
		WithPartitions(20).
		WithReplicas(3).
		WithTopicConfiguration(kafka.NewTopicConfiguration().
			WithCleanupPolicy([]string{"delete", "compact"}).
			WithRetentionMs(604800000).
			WithRetentionBytes(1000000000).
			WithDeleteRetentionMs(86400000).
			WithMaxMessageBytes(1048588)).
		WithBindingVersion("0.5.0")

	// Create a channel and add the Kafka binding.
	channel := asyncapi2.NewChannel().
		WithAddress("user/signup").
		WithPublish(asyncapi2.NewOperation().
			WithMessage(asyncapi2.NewMessage().
				WithPayload(map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"userId": map[string]interface{}{
							"type": "string",
						},
					},
				}))).
		WithBindings(map[string]interface{}{
			"kafka": kafkaBinding,
		})

	// Add the channel to the document.
	doc = doc.WithChannel("user/signup", channel)

	// Validate the document (optional but recommended).
	if err := doc.Validate(); err != nil {
		log.Fatalf("Validation error: %v", err)
	}

	// Convert the document to JSON (optional).
	jsonData, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	fmt.Printf("AsyncAPI Document with Kafka Binding:\n%s\n", string(jsonData))
}
```
