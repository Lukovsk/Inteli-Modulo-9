package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type SensorData struct {
	idSensor     string
	timestamp    time.Time
	tipoPoluente string
	nivel        float32
}

func Producer() *kafka.Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092, localhost:39092",
		"client.id":         "go-producer",
	})

	if err != nil {
		panic(err)
	}

	defer producer.Close()
	return producer
}

func Consumer() *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092, localhost: 39092",
		"group.id":          "go-consumer-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	return consumer
}

func TestIntegrity(t *testing.T) {
	producer := Producer()

	topic := "test"
	message := fmt.Sprintf(`{%v}`, SensorData{idSensor: "test-sensor", timestamp: time.Now(), tipoPoluente: "test", nivel: 1})
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	producer.Flush(10 * 1000)

	consumer := Consumer()

	consumer.SubscribeTopics([]string{topic}, nil)
	fmt.Println("Conected!")

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			if string(msg.Value) == message {
				t.Log("Integration succeeded!")
			}

		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

}
