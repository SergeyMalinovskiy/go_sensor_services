package main

import (
	"github.com/SergeyMalinovskiy/growther/internal/app/server/routes"
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()

	routes.SetupRouting(serveMux)

	log.Println("Server starting...")
	err := http.ListenAndServe(":8080", serveMux)
	if err != nil {
		log.Fatalf("Server running error: %v", err)
	}
}
