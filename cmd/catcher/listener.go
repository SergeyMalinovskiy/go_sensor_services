package main

import (
	"fmt"
	sensorServiceContract "github.com/SergeyMalinovskiy/growther/libs/contract/golang"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

type MessageHandler func(message ContainsMessage) int

type CatcherListener interface {
	Listen()
	AddHandler(channel string, handler MessageHandler)
}

type ContainsMessage interface {
	GetChannel() string
	GetMessage() sensorServiceContract.EventPayload
}

type MqttListener struct {
	mqttClient mqtt.Client
}

func NewMqttListener(client mqtt.Client) MqttListener {
	listener := MqttListener{
		mqttClient: client,
	}

	if !listener.mqttClient.IsConnected() {
		listener.connect()
	}

	return listener
}

func (listener MqttListener) Listen() {
	defer func(mqttClient *mqtt.Client) {
		fmt.Println("MQTT: Disconnecting...")
		(*mqttClient).Disconnect(1)
	}(&listener.mqttClient)

	fmt.Println("MQTT: Listening...")

	select {}
}

func (listener MqttListener) AddHandler(channel string, handler MessageHandler) {
	messageHandler := func(client mqtt.Client, message mqtt.Message) {
		code := handler(BusEvent{channel: message.Topic(), message: message.Payload()})

		fmt.Printf("MQTT: handler result code %d\n", code)
	}

	if token := listener.mqttClient.Subscribe(channel, 0, messageHandler); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT: Subscribe channel error: %v", token.Error())
	}

	fmt.Printf("MQTT: Subscribed to topic \"%s\"\n", channel)
}

func (listener MqttListener) connect() {
	fmt.Println("MQTT: Connecting...")
	if token := listener.mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}
