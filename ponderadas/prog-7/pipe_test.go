package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	godotenv "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestPipeline(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading environment: %s", err)
	}

	MONGO_URL := os.Getenv("MONGO_URL")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGO_URL).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	database := client.Database("pond")
	colection := database.Collection("Sensors")
	var topic = "/mqtt/sensor"

	data := Sensor{
		NH3_ppm: rand.Intn(400),
		CO_ppm:  rand.Intn(1000),
		NO2_ppm: rand.Intn(30),
	}

	publisher := Publisher()

	jsonData, _ := json.Marshal(data)

	publisher.Publish(topic, 1, false, jsonData)

	time.Sleep(5 * time.Second)

	filter := bson.D{{Key: "NH3_ppm", Value: data.NH3_ppm}, {Key: "CO_ppm", Value: data.CO_ppm}, {Key: "NO2_ppm", Value: data.NO2_ppm}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := colection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	var mongoData []Sensor
	if err = cursor.All(ctx, &mongoData); err != nil {
		panic(err)
	}
	result := mongoData[0]

	if result.sensor == data.sensor && result.CO_ppm == data.CO_ppm && result.NO2_ppm == data.NO2_ppm && result.NH3_ppm == data.NH3_ppm {
		t.Log("Passed")
	}

}
