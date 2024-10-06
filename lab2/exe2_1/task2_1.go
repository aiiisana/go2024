package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	connStr := "user=aiiisana password=mypassword dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(0)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
	return db
}

func createTable(db *sql.DB) {
	query := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name TEXT UNIQUE NOT NULL,
            age INT NOT NULL
        );
    `
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
	fmt.Println("Table 'users' created successfully!")
}

func insertUsers(db *sql.DB, users []User) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Error starting transaction:", err)
	}

	stmt, err := tx.Prepare("INSERT INTO users (name, age) VALUES ($1, $2)")
	if err != nil {
		tx.Rollback()
		log.Fatal("Error preparing statement:", err)
	}
	defer stmt.Close()

	for _, user := range users {
		_, err := stmt.Exec(user.Name, user.Age)
		if err != nil {
			tx.Rollback()
			log.Fatalf("Error inserting user %s: %v", user.Name, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction:", err)
	}

	fmt.Println("Users inserted successfully!")
}

func queryUsers(db *sql.DB, minAge int, page, pageSize int) {
	offset := (page - 1) * pageSize

	query := `
        SELECT id, name, age FROM users
        WHERE age >= $1
        ORDER BY id
        LIMIT $2 OFFSET $3
    `
	rows, err := db.Query(query, minAge, pageSize, offset)
	if err != nil {
		log.Fatal("Error querying users:", err)
	}
	defer rows.Close()

	fmt.Printf("Page %d results:\n", page)
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal("Error in row iteration:", err)
	}
}

func updateUser(db *sql.DB, id int, name string, age int) {
	query := `UPDATE users SET name = $1, age = $2 WHERE id = $3`
	result, err := db.Exec(query, name, age, id)
	if err != nil {
		log.Fatal("Error updating user:", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Error retrieving affected rows:", err)
	}

	if rowsAffected == 0 {
		fmt.Printf("No user with ID %d found.\n", id)
	} else {
		fmt.Printf("User with ID %d updated successfully.\n", id)
	}
}

func deleteUser(db *sql.DB, id int) {
	query := `DELETE FROM users WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatal("Error deleting user:", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal("Error retrieving affected rows:", err)
	}

	if rowsAffected == 0 {
		fmt.Printf("No user with ID %d found.\n", id)
	} else {
		fmt.Printf("User with ID %d deleted successfully.\n", id)
	}
}

type User struct {
	Name string
	Age  int
}

func main() {
	db := connectDB()
	defer db.Close()

	createTable(db)

	users := []User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	insertUsers(db, users)

	queryUsers(db, 20, 1, 2)

	updateUser(db, 1, "Alice Updated", 31)

	deleteUser(db, 2)

	queryUsers(db, 20, 1, 2)
}
