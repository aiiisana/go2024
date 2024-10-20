package delivery

import (
	"taskmanager/internal/repository"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *mux.Router, db *gorm.DB) {
	userRoutes := router.PathPrefix("/auth").Subrouter()
	userRoutes.HandleFunc("/users", repository.RegisterUser(db)).Methods("POST")
	userRoutes.HandleFunc("/login", repository.LoginUser(db)).Methods("POST")

	userRoutes.HandleFunc("/users", repository.GetUsers(db)).Methods("GET")
}
