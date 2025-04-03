package main

import sensorServiceContract "github.com/SergeyMalinovskiy/growther/libs/contract/golang"

type HandlersResolver struct {
	dataHandler   sensorServiceContract.DataHandler
	readyHandler  sensorServiceContract.ReadyHandler
	brokenHandler sensorServiceContract.BrokenHandler
}

func NewHandlersResolver(
	dataHandler sensorServiceContract.DataHandler,
	readyHandler sensorServiceContract.ReadyHandler,
	brokenHandler sensorServiceContract.BrokenHandler,
) HandlersResolver {
	return HandlersResolver{
		dataHandler:   dataHandler,
		readyHandler:  readyHandler,
		brokenHandler: brokenHandler,
	}
}

func (catcher HandlersResolver) GetDataHandler() sensorServiceContract.DataHandler {
	return catcher.dataHandler
}

func (catcher HandlersResolver) GetReadyHandler() sensorServiceContract.ReadyHandler {
	return catcher.readyHandler
}

func (catcher HandlersResolver) GetBrokenHandler() sensorServiceContract.BrokenHandler {
	return catcher.brokenHandler
}
