package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	godotenv "github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	var API_URL = os.Getenv("MONGO_URL")

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(API_URL).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	database := client.Database("Modulo9")

	colection := database.Collection("Sensors")

	// go Publisher()
	Subscriber(colection)
	select {}
}

type Sensor struct {
	NH3_ppm, CO_ppm, NO2_ppm int
	sensor                   string
	// timestamp                time.Time
}

func Insert(data Sensor, db *mongo.Collection) {

	document := bson.D{{Key: "sensor", Value: data.sensor}, {Key: "NH3_ppm", Value: data.NH3_ppm}, {Key: "CO_ppm", Value: data.CO_ppm}, {Key: "NO2_ppm", Value: data.NO2_ppm}}
	response, err := db.InsertOne(context.TODO(), document)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Document %s inserted successfully, id: %s", document, response.InsertedID)
}

func Select(db *sql.DB) {
	row, err := db.Query("SELECT * FROM sensor ORDER BY timestamp")
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer row.Close()
	for row.Next() {
		var id int
		var sensor string
		var NH3_ppm int
		var CO_ppm int
		var NO2_ppm int
		var timestamp time.Time
		row.Scan(&id, &sensor, &NH3_ppm, &CO_ppm, &NO2_ppm, &timestamp)
		log.Println("Sensor data: %v - %v - %v - %v - %v - %v", id, sensor, NH3_ppm, CO_ppm, NO2_ppm, timestamp)
	}
}
