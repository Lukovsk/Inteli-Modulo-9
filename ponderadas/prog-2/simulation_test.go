package main

import (
	"encoding/json"
	"testing"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func TestConection(t *testing.T) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("test-client")

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Errorf("Error in connection with broker MQTT: %v", token.Error())
	} else {
		t.Log("Connection with broker MQTT succeeded")
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
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("test-client")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to connect MQTT broker: %v", token.Error())
	}

	received := make(chan bool)
	token := client.Subscribe("/pond-2", 1, func(client MQTT.Client, msg MQTT.Message) {
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

	msg := SensorData()
	jsonData, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("Error to convert to JSON: %v", err)
	}

	data := string(jsonData)
	token = client.Publish("/pond-2", 0, false, data)
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to publish message: %v", token.Error())
	}

	select {
	case <-received:
		t.Log("Message received")
	case <-time.After(5 * time.Second):
		t.Fatalf("Timeout")
	}
}
