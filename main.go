package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func newPostgreClient() *sql.DB {
	// All of the parameters needed to connect to Postgre.
	pgUser := os.Getenv("PGUSER")
	// You should have a password! This is for testing only.
	// pgPassword := os.Getenv("PGPASSWORD")
	pgHost := os.Getenv("PGHOST")
	pgPort := os.Getenv("PGPORT")
	pgDatabase := os.Getenv("PGDATABASE")
	pgSSLMode := os.Getenv("PGSSLMODE")

	connStr := fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=%s",
		pgUser, pgDatabase, pgHost, pgPort, pgSSLMode)
	fmt.Println("Postgres connection string:", connStr)

	// Create a postgres db connection.
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return conn
}

// A helper method to print the rows in the users table.
func printRows(rows *sql.Rows) {
	defer rows.Close()
	for rows.Next() {
		var id int
		var lastname string
		var firstname string

		err := rows.Scan(&id, &lastname, &firstname)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, lastname, firstname)
	}
}

func main() {
	fmt.Println("Service started.")

	db := newPostgreClient()
	defer db.Close()

	// Verify DB connection.
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to Postgres DB")

	// Some examples of issuing SQL queries to Postgres.
	// Create table
	result, _ := db.Exec("CREATE TABLE users (ID int, lastname varchar(255), firstname varchar(255));")
	fmt.Println(result)

	// Insert rows into table
	result, _ = db.Exec("INSERT INTO users(id, lastname, firstname) VALUES (1, 'A', 'B');")
	fmt.Println(result)
	result, _ = db.Exec("INSERT INTO users(id, lastname, firstname) VALUES (2, 'C', 'D');")
	fmt.Println(result)

	// Select all.
	rows, _ := db.Query("SELECT * FROM users;")
	printRows(rows)

	// Update row.
	result, _ = db.Exec("UPDATE users SET lastname='E', firstname='F' where id=2;")
	rows, _ = db.Query("SELECT * FROM users;")
	printRows(rows)

	// Delete row.
	result, _ = db.Exec("DELETE FROM users where id=2;")
	rows, _ = db.Query("SELECT * FROM users;")
	printRows(rows)

	// Create index on id column.
	result, _ = db.Exec("CREATE UNIQUE INDEX index_id ON users (id);")
	fmt.Println(result)
}
