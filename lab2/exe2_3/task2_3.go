package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func initDB() (*sql.DB, error) {
	connStr := "user=aiiisana password=mypassword dbname=mydb sslmode=disable"
	return sql.Open("postgres", connStr)
}

func initGORM() (*gorm.DB, error) {
	dsn := "user=aiiisana password=mypassword dbname=mydb sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func getUsersSQL(w http.ResponseWriter, r *http.Request) {
	db, err := initDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	age := r.URL.Query().Get("age")
	sort := r.URL.Query().Get("sort")

	query := "SELECT id, name, age FROM users"
	if age != "" {
		query += " WHERE age = " + age
	}
	if sort == "name" {
		query += " ORDER BY name"
	}

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func createUserSQL(w http.ResponseWriter, r *http.Request) {
	db, err := initDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO users (name, age) VALUES ($1, $2)`
	_, err = db.Exec(query, user.Name, user.Age)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func updateUserSQL(w http.ResponseWriter, r *http.Request) {
	db, err := initDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE users SET name = $1, age = $2 WHERE id = $3`
	_, err = db.Exec(query, user.Name, user.Age, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteUserSQL(w http.ResponseWriter, r *http.Request) {
	db, err := initDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := `DELETE FROM users WHERE id = $1`
	_, err = db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getUsersGORM(w http.ResponseWriter, r *http.Request) {
	db, err := initGORM()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var users []User
	query := db

	age := r.URL.Query().Get("age")
	if age != "" {
		ageInt, _ := strconv.Atoi(age)
		query = query.Where("age = ?", ageInt)
	}

	sort := r.URL.Query().Get("sort")
	if sort == "name" {
		query = query.Order("name")
	}

	if err := query.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func createUserGORM(w http.ResponseWriter, r *http.Request) {
	db, err := initGORM()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func updateUserGORM(w http.ResponseWriter, r *http.Request) {
	db, err := initGORM()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.Model(&User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteUserGORM(w http.ResponseWriter, r *http.Request) {
	db, err := initGORM()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := mux.Vars(r)
	id := params["id"]

	if err := db.Delete(&User{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", getUsersSQL).Methods("GET")
	r.HandleFunc("/users", createUserSQL).Methods("POST")
	r.HandleFunc("/users/{id}", updateUserSQL).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUserSQL).Methods("DELETE")

	r.HandleFunc("/gorm/users", getUsersGORM).Methods("GET")
	r.HandleFunc("/gorm/users", createUserGORM).Methods("POST")
	r.HandleFunc("/gorm/users/{id}", updateUserGORM).Methods("PUT")
	r.HandleFunc("/gorm/users/{id}", deleteUserGORM).Methods("DELETE")

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
