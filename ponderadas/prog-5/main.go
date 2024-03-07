package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, _ := sql.Open("sqlite3", "./db/db.db")
	defer db.Close()

	createTables(database)

	// test_data := Sensor{sensor: "test", CO_ppm: 0, NH3_ppm: 0, NO2_ppm: 0}
	// Insert(test_data, db)

	go Publisher()
	Subscriber(database)
	select {}
}

func createTables(db *sql.DB) {
	sensorTableStmt := `
	CREATE TABLE IF NOT EXISTS sensor (id INTEGER PRIMARY KEY, sensor TEXT, NH3_ppm INTEGER, CO_ppm INTEGER, NO2_ppm INTEGER)`
	// userTableStmt := `CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY, name TEXT, password TEXT)`

	command, err := db.Prepare(sensorTableStmt)
	if err != nil {
		log.Fatal(err.Error())
	}
	command.Exec()

	// command2, err := db.Prepare(userTableStmt)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// command2.Exec()
}

type Sensor struct {
	NH3_ppm, CO_ppm, NO2_ppm int
	sensor                   string
	// timestamp                time.Time
}

func Insert(data Sensor, db *sql.DB) {
	// stmt := `INSERT INTO sensor(sensor, NH3_ppm, CO_ppm, NO2_ppm) VALUES (?, ?, ?, ?)`

	// statement, err := db.Prepare(stmt)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// _, err = statement.Exec(data.sensor, data.NH3_ppm, data.CO_ppm, data.NO2_ppm)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
	_, err := db.Exec(fmt.Sprintf("INSERT INTO sensor (sensor, NH3_ppm, CO_ppm, NO2_ppm) VALUES ('%v', %v, %v, %v)", data.sensor, data.NH3_ppm, data.CO_ppm, data.NO2_ppm))
	if err != nil {
		log.Fatal(err)
	}
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
