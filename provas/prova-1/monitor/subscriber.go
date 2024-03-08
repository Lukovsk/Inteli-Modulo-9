package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

type Sensor struct {
	Id        string
	Group     string
	Temp      int32
	Timestamp string
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

var messageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("[SUBSCRIBER][%s] %s \n", msg.Topic(), msg.Payload())
	result := strings.Split(string(msg.Payload()), ",")
	fmt.Printf("%v \n", result)

	id := strings.Split(strings.Split(result[0], ":")[1], "-")[1]
	group := strings.Split(result[1], ":")[1]
	temp := strings.Split(result[2], ":")[1]
	times := strings.Split(result[3], `":"`)[1]
	var alert = ""
	if group == "freezer" {
		if strconv.Atoi(temp) > -15 {
			alert = ""
		}
	}

	fmt.Printf("Lj %v: %v %v | %v %v", id, group, temp, times)

}

func Subscriber() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading environment: %s", err)
	}

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	var topic = "/prova/temperature"

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("Subscriber")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.SetDefaultPublishHandler(messageSubHandler)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	select {}
}

func main() {
	Subscriber()
}
