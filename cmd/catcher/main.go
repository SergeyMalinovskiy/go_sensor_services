package main

import (
	"database/sql"
	"fmt"
	"github.com/SergeyMalinovskiy/growther/cmd/catcher/repositories"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = "5439"
	user   = "growth_monitor"
	pass   = "growth_monitor"
	dbname = "growth_monitor"
)

func main() {
	fmt.Println("Catcher init...")

	loadConfig()
	db := configureDatabase()
	defer dbClose(db)

	runWatchProcess(db)
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file. %v", err.Error())
	}
}

func configureDatabase() *sql.DB {
	pqSqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		pass,
		user,
		dbname,
	)

	db, err := sql.Open("postgres", pqSqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func dbClose(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func runWatchProcess(db *sql.DB) {
	host := os.Getenv("MQTT_BROKER_HOST")
	port := os.Getenv("MQTT_BROKER_PORT")

	mqttClient := mqttClientInit(host, port)

	listener := NewMqttListener(
		mqttClient,
	)

	catcher := NewCatcher(
		listener,
		repositories.NewMeasurementRepository(db),
		repositories.NewSensorRepository(db),
	)

	catcher.RegisterHandlers()

	catcher.RunCatch()
}
