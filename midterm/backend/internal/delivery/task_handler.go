package delivery

import (
	"taskmanager/internal/repository"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetupTaskRoutes(router *mux.Router, db *gorm.DB) {
	taskRoutes := router.PathPrefix("/users/{userId}/tasks").Subrouter()
	taskRoutes.HandleFunc("/", repository.GetTasks(db)).Methods("GET")
	taskRoutes.HandleFunc("/", repository.CreateTask(db)).Methods("POST")
	taskRoutes.HandleFunc("/{id}", repository.UpdateTask(db)).Methods("PUT")
	taskRoutes.HandleFunc("/{id}", repository.DeleteTask(db)).Methods("DELETE")
}
