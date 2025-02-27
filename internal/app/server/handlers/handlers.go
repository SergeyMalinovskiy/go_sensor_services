package handlers

import (
	"fmt"
	"net/http"
)

func NotFoundHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Not found!")
}

func WelcomeHandler(writer http.ResponseWriter, request *http.Request) {
	var congratulation = "Hello my dear friend:)"
	fmt.Println("Welcome request!")

	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte(congratulation))

	if err != nil {
		fmt.Printf("Error %v", err)
	}
}
