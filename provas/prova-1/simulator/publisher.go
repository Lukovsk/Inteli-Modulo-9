package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

type Sensor struct {
	Id        string
	Group     string
	Temp      int32
	Timestamp string
}

func GetData() *Sensor {
	lj := rand.Intn(4)

	var group [2]string
	if lj%2 == 0 {
		group[0] = "f"
		group[1] = "freezer"
	} else {
		group[0] = "g"
		group[1] = "geladeira"
	}

	id := "lj-" + strconv.Itoa(lj) + "-" + group[0] + "-" + strconv.Itoa(rand.Intn(4))

	data := &Sensor{
		Id:        id,
		Group:     group[1],
		Temp:      rand.Int31n(45) - 30,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	return data
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
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

	topic := "/prova/temperature"

	for {
		data := GetData()

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting data to JSON", err)
			return
		}

		token := client.Publish(topic, 1, false, string(jsonData))
		token.Wait()

		fmt.Println("[PUBLISHER][" + topic + "] " + string(jsonData))
		time.Sleep(2 * time.Second)
	}
}

func main() {
	Publisher()
}
