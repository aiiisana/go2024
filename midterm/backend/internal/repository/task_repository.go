package repository

import (
	"encoding/json"
	"net/http"
	"strconv"
	"taskmanager/internal/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateTask(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["userId"])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		task.UserID = uint(userID)

		if err := db.Create(&task).Error; err != nil {
			http.Error(w, "Could not create task", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
	}
}

func GetTasks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["userId"])
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		var tasks []models.Task
		if err := db.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
			http.Error(w, "Could not retrieve tasks", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(tasks)
	}
}

func GetTask(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task
		userID := r.Context().Value("user_id").(uint)

		db.Where("user_id = ?", userID).Find(&task)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	}
}

func UpdateTask(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task
		id := mux.Vars(r)["id"]
		if err := db.Where("id = ?", id).First(&task).Error; err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		db.Save(&task)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(task)
	}
}

func DeleteTask(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task
		id := mux.Vars(r)["id"]
		if err := db.Where("id = ?", id).First(&task).Error; err != nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		db.Delete(&task)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Task deleted successfully")
	}
}
