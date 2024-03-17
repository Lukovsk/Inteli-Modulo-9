package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

// var data Sensor

var db *mongo.Collection


var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

var messageSubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("[SUBSCRIBER][%s] %s \n", msg.Topic(), msg.Payload())
	result := strings.Split(string(msg.Payload()), " - ")

	// timestamp, _ := time.Parse(result[0], "2006-01-02T15:04:05.000Z")
	name := result[1]
	nh3, _ := strconv.Atoi(result[2])
	co, _ := strconv.Atoi(result[3])
	no2, _ := strconv.Atoi(result[4])

	data := Sensor{NH3_ppm: nh3, CO_ppm: co, NO2_ppm: no2, sensor: name}
	Insert(data, db)
}

func Subscriber(dbPointer *mongo.Collection) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading environment: %s", err)
	}

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	var topic = "/mqtt/sensor"

		 = dbPointer

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
