package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var defaultMessageHandler mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("mqtt: Received message: %s from topic: %s\n", message.Payload(), message.Topic())
}

func main() {
	fmt.Println("Catcher init...")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file. %v", err.Error())
	}

	host := os.Getenv("MQTT_BROKER_HOST")
	port := os.Getenv("MQTT_BROKER_PORT")

	mqttClient := mqttClientInit(host, port)

	fmt.Println("Connecting...")
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	select {}
}

func mqttClientInit(host string, port string) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(fmt.Sprintf("tcp://%s:%s", host, port))

	opts.SetClientID("go_mqtt_client_id")
	opts.SetDefaultPublishHandler(defaultMessageHandler)

	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connected to MQTT broker")

		if token := client.Subscribe("test/#", 0, nil); token.Wait() && token.Error() != nil {
			log.Fatalf("Subscribe channel error: %v", token.Error())
		}

		fmt.Println("Subscribed to topic: \"test/#\"")
	}

	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v\n", err)
	}

	return mqtt.NewClient(opts)
}
