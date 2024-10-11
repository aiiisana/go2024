package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func createTable(db *sql.DB) {
	query := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name TEXT,
            age INT
        );
    `
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func insertUsers(db *sql.DB, name string, age int) {
	query := `insert into users (name, age) values ($1, $2)`
	_, err := db.Exec(query, name, age)
	if err != nil {
		panic(err)
	}
}

func printUsers(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}

func main() {
	connStr := "user=aiiisana dbname=mydb password=mypassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTable(db)
	// insertUsers(db, "John Doe", 30)
	//insertUsers(db, "Aisana", 18)
	printUsers(db)
}
