package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SergeyMalinovskiy/growther/cmd/catcher/repositories"
	sensorservicecontract "github.com/SergeyMalinovskiy/growther/libs/contract/golang"
	"log"
	"regexp"
	"strconv"
	"time"
)

type MeasurementRawData struct {
	Value      string
	ValueExtra string
	SensorId   int
	Datetime   string
}

type Catcher struct {
	listener           CatcherListener
	measurementSaver   repositories.MeasurementSaver
	sensorsInformation repositories.CanReceiveSensorData
}

func NewCatcher(
	listener CatcherListener,
	measurementSaver repositories.MeasurementSaver,
	sensorsInformation repositories.CanReceiveSensorData,
) Catcher {
	return Catcher{
		listener: listener,
		//handlersResolver: nil, // TODO:а этот модуль тут нужен вообще?
		measurementSaver:   measurementSaver,
		sensorsInformation: sensorsInformation,
	}
}

func (c Catcher) RegisterHandlers() {
	dataMaskRe := regexp.MustCompile(`^test/sensors/([^/]+)/data$`)

	c.listener.AddHandler("test/sensors/+/data", func(message ContainsMessage) int {
		isValid, sensorId := c.validateChannel(message.GetChannel(), dataMaskRe)
		if !isValid {
			log.Fatalf("Catcher: sensor id is invalid!")
		}

		fmt.Printf("\nCatcher: Sensor ID: %v;\nDATA %s\n", sensorId, message.GetMessage())
		measurementRawData := MeasurementRawData{}
		err := json.Unmarshal(message.GetMessage(), &measurementRawData)
		if err != nil {
			fmt.Printf("Catcher: parsing error: %v\n", err)

			return 0
		}

		measurementRawData.SensorId = sensorId

		measurementData, err := c.prepareMeasurementData(measurementRawData)

		if err != nil {
			return 0
		}

		_, err = c.measurementSaver.Save(measurementData)
		if err != nil {
			fmt.Printf("Catcher: saving data error: %v\n", err)
		}

		return 1
	})

	c.listener.AddHandler("test/sensors/+/ready", func(message ContainsMessage) int {
		fmt.Printf("\nTOPIC: %s;\nDATA: %s;\n", message.GetChannel(), message.GetMessage())

		return 1
	})

	c.listener.AddHandler("test/sensors/+/broken", func(message ContainsMessage) int {
		fmt.Printf("\nTOPIC: %s;\nDATA: %s;\n", message.GetChannel(), message.GetMessage())

		return 1
	})

	//client.Subscribe("test/sensors/#/ready", 0, func(client mqtt.Client, message mqtt.EventPayload) {

	//})
	//
	//client.Subscribe("test/sensors/#/broken", 0, func(client mqtt.Client, message mqtt.EventPayload) {

	//})
}

func (c Catcher) RunCatch() {
	c.listener.Listen()
}

func (c Catcher) parseSensorIdFromChannel(channel string, maskRegexp *regexp.Regexp) (string, error) {
	matches := maskRegexp.FindStringSubmatch(channel)

	if len(matches) > 1 {
		return matches[1], nil
	}

	return "", errors.New(fmt.Sprintf("sensor id parsing error: \"%s\"\n", channel))
}

func (c Catcher) validateChannel(channel string, channelMaskRe *regexp.Regexp) (bool, int) {
	sensorIdChannel, err := c.parseSensorIdFromChannel(channel, channelMaskRe)
	if err != nil {
		fmt.Printf("Catcher: sensor validate error: %v\n", err)
		return false, -1
	}

	sensorId, err := strconv.Atoi(sensorIdChannel)
	if err != nil {
		fmt.Printf("Catcher: sensor validate error: %v\n", err)
		return false, -1
	}

	sensorExists := c.sensorsInformation.Exists(sensorId)
	if !sensorExists {
		fmt.Printf("Catcher: sensor validate error: sensor with id = \"%d\" not exists.\n", sensorId)
		return false, -1
	}

	return true, sensorId
}

func (c Catcher) prepareMeasurementData(rawData MeasurementRawData) (sensorservicecontract.MeasurementData, error) {
	isValid := true

	value, err := strconv.ParseFloat(rawData.Value, 32)
	if err != nil {
		fmt.Printf("Catcher: raw data value field validation error: %v\n", err)
		isValid = false
	}

	datetime, err := time.Parse("2006-01-02 15:04:05", rawData.Datetime)
	if err != nil {
		fmt.Printf("Catcher: raw data datetime field validation error: %v\n", err)
		isValid = false
	}

	if !isValid {
		return sensorservicecontract.MeasurementData{}, errors.New("catcher: prepare measurement data error")
	}

	data := sensorservicecontract.MeasurementData{
		Value:      float32(value),
		ValueExtra: rawData.ValueExtra,
		Datetime:   datetime,
		SensorId:   rawData.SensorId,
	}

	return data, nil
}
