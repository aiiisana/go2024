package main

import (
	"log"
	"net/http"
	"taskmanager/internal/config"
	"taskmanager/internal/delivery"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db := config.SetupDatabase()

	router := mux.NewRouter()

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	router.Use(corsMiddleware)

	delivery.SetupTaskRoutes(router, db)
	delivery.SetupUserRoutes(router, db)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(router)))
}
