package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

func SensorData() map[string]int {
	data := map[string]int{
		"NH3_ppm": rand.Intn(400),
		"CO_ppm":  rand.Intn(1000),
		"NO2_ppm": rand.Intn(30),
	}
	return data
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Recebido: %s do tópico: %s com QoS: %d\n", msg.Payload(), msg.Topic(), msg.Qos())
}

func Publisher() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("Publisher")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
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

		topic := "/mqtt/sensor"

		msg := time.Now().Format(time.RFC3339) + " " + string(jsonData)

		token := client.Publish(topic, 1, false, msg)
		token.Wait()

		fmt.Println("[PUBLISHER][" + topic + "] " + msg)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	Publisher()
}
