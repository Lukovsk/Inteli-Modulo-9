package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type SensorData struct {
	idSensor     string
	timestamp    time.Time
	tipoPoluente string
	nivel        float32
}

func generateSensorData(tipo string) SensorData {
	timenow := time.Now()
	return SensorData{
		idSensor:     "sensor-1",
		timestamp:    timenow,
		tipoPoluente: tipo,
		nivel:        float32(rand.Intn(200) / 12),
	}
}

func generateALotOfData(much int) []SensorData {
	data := []SensorData{}
	for i := 0; i < much; i++ {
		data = append(data, generateSensorData("PM2.5"))
		data = append(data, generateSensorData("PM5"))
		data = append(data, generateSensorData("PM10"))
	}
	return data
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

func main() {
	producer := Producer()
	fmt.Println("Conected!")

	topic := "air"
	message := []byte(fmt.Sprintf(`%v`, generateALotOfData(4)))
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)
	fmt.Printf("Produced: %s", message)

	producer.Flush(10 * 1000)

	select {}
}
