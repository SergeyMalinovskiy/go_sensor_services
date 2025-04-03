package main

import sensorServiceContract "github.com/SergeyMalinovskiy/growther/libs/contract/golang"

type BusEvent struct {
	channel string
	message sensorServiceContract.EventPayload
}

func (event BusEvent) GetChannel() string {
	return event.channel
}

func (event BusEvent) GetMessage() sensorServiceContract.EventPayload {
	return event.message
}
