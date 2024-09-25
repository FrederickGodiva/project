package main

import (
	"backend/config"
	"backend/routes"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func main() {
	database.Connect()

	r := routes.SetupRouter()

	corsObj := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj, corsMethods, corsHeaders)(r)))
}
