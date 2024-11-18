package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/health", healthCheckHandler)

	fmt.Printf("Server listening and running on 8048 port...\n\r")
	err := http.ListenAndServe(":8048", nil)
	if err != nil {
		fmt.Printf("Err: %s", err.Error())
	}
}

func healthCheckHandler(writer http.ResponseWriter, request *http.Request) {
	testData := "I'm health!"

	writer.WriteHeader(200)

	_, err := writer.Write([]byte(testData))
	if err != nil {
		fmt.Printf("Err: %s", err.Error())
	}
}

func handlerFunc(writer http.ResponseWriter, r *http.Request) {
	testData := "OK!"

	writer.WriteHeader(200)

	_, err := writer.Write([]byte(testData))
	if err != nil {
		fmt.Printf("Err: %s", err.Error())
	}
}
