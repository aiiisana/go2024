package repository

import (
	"encoding/json"
	"net/http"
	"taskmanager/internal/models"

	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := db.Create(&user).Error; err != nil {
			http.Error(w, "Could not create user", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("User registered successfully")
	}
}

func LoginUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		var user models.User
		if err := db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		if user.Password != loginData.Password {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		response := struct {
			UserId uint `json:"userId"`
		}{
			UserId: user.ID,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func GetUsers(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []models.User

		if err := db.Find(&users).Error; err != nil {
			http.Error(w, "Could not retrieve users", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}
