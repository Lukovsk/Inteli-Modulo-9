package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type SensorData struct {
	idSensor     string    `json:"idSensor"`
	timestamp    time.Time `json:"timestamp"`
	tipoPoluente string    `json:"tipoPoluente"`
	nivel        float32   `json:"nivel"`
}

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092, localhost: 39092",
		"group.id":          "go-consumer-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	topic := "air"
	consumer.SubscribeTopics([]string{topic}, nil)
	fmt.Println("Conected!")

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
}
