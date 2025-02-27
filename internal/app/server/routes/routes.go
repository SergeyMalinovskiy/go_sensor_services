package routes

import (
	"github.com/SergeyMalinovskiy/growther/internal/app/server/handlers"
	"net/http"
)

func SetupRouting(serveMux *http.ServeMux) {
	serveMux.HandleFunc("/", handlers.WelcomeHandler)
}
