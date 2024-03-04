package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

func TestDotenv(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	var broker = os.Getenv("BROKER_ADDR")
	fmt.Printf("Broker address: %v \n", broker)
	var username = os.Getenv("HIVE_USER")
	fmt.Printf("Username: %v \n", username)
	var password = os.Getenv("HIVE_PSWD")
	fmt.Printf("Password: %v \n", password)
}

func TestConection(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883

	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("Publisher")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

func TestDataValidation(t *testing.T) {
	msg := SensorData()

	expectedFields := []string{"NH3_ppm", "CO_ppm", "NO2_ppm"}
	for _, field := range expectedFields {
		if _, ok := msg[field]; !ok {
			t.Errorf("Expected field: %s", field)
			return
		}
	}
	t.Log("Data validation successfull")
}

func TestPublisher(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883

	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("Publisher")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to connect MQTT broker: %v", token.Error())
	}

	topic := "/mqtt/sensor"

	received := make(chan bool)
	token := client.Subscribe(topic, 1, func(client MQTT.Client, msg MQTT.Message) {
		// Validating message
		var data map[string]int
		if err := json.Unmarshal(msg.Payload(), &data); err != nil {
			t.Errorf("Error validating message: %v", err)
			return
		}

		// Fields validation
		expectedFields := []string{"NH3_ppm", "CO_ppm", "NO2_ppm"}
		for _, field := range expectedFields {
			if _, ok := data[field]; !ok {
				t.Errorf("Field %s expected but not received", field)
				return
			}
		}

		received <- true
	})
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to subscribe MQTT topic: %v", token.Error())
	}

	data := SensorData()
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Error to convert to JSON: %v", err)
	}

	msg := string(jsonData)
	token = client.Publish(topic, 0, false, msg)
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to publish message: %v", token.Error())
	}
	token.Wait()

	select {
	case <-received:
		t.Log("Message received")
	case <-time.After(5 * time.Second):
		t.Fatalf("Timeout")
	}
}
