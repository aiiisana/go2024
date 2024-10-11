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

// GORM
type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB
var gormDB *gorm.DB

func init() {
	var err error
	connStr := "user=aiiisana password=mypassword dbname=mydb sslmode=disable"

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database via GORM:", err)
	}

	gormDB.AutoMigrate(&User{})
}

func getUsersSQL(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

func getUsersGORM(w http.ResponseWriter, r *http.Request) {
	var users []User
	result := gormDB.Find(&users)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func createUserSQL(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	query := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, user.Name, user.Age).Scan(&user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func createUserGORM(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	result := gormDB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func updateUserSQL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	query := `UPDATE users SET name=$1, age=$2 WHERE id=$3`
	_, err := db.Exec(query, user.Name, user.Age, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ID = uint(id)
	json.NewEncoder(w).Encode(user)
}

func updateUserGORM(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	result := gormDB.Model(&User{}).Where("id = ?", id).Updates(User{Name: user.Name, Age: user.Age})
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	user.ID = uint(id)
	json.NewEncoder(w).Encode(user)
}

func deleteUserSQL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	query := `DELETE FROM users WHERE id=$1`
	_, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User with ID %d deleted", id)
}

func deleteUserGORM(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	result := gormDB.Delete(&User{}, id)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User with ID %d deleted", id)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/sql/users", getUsersSQL).Methods("GET")
	router.HandleFunc("/sql/user", createUserSQL).Methods("POST")
	router.HandleFunc("/sql/user/{id}", updateUserSQL).Methods("PUT")
	router.HandleFunc("/sql/user/{id}", deleteUserSQL).Methods("DELETE")

	router.HandleFunc("/gorm/users", getUsersGORM).Methods("GET")
	router.HandleFunc("/gorm/user", createUserGORM).Methods("POST")
	router.HandleFunc("/gorm/user/{id}", updateUserGORM).Methods("PUT")
	router.HandleFunc("/gorm/user/{id}", deleteUserGORM).Methods("DELETE")

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
