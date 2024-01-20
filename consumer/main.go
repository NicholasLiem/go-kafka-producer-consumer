package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	err = c.SubscribeTopics([]string{"test_topics"}, nil)
	if err != nil {
		fmt.Printf("Error subscribing to topic: %v\n", err)
	} else {
		fmt.Println("Subscribed to topic: test_topics")
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
}
