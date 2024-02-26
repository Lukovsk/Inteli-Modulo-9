package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func SensorData() map[string]int {
	data := map[string]int{
		"NH3_ppm": rand.Intn(400),
		"CO_ppm":  rand.Intn(1000),
		"NO2_ppm": rand.Intn(30),
	}
	return data
}

func Client() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("publisher")

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	for {
		data := SensorData()
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting data to JSON", err)
			return
		}

		msg := time.Now().Format(time.RFC3339) + " " + string(jsonData)

		token := client.Publish("/pond-2", 1, false, msg)
		token.Wait()

		fmt.Println("[PUBLISHER] ", msg)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	Client()
}
