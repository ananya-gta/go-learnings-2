package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

// declare a global variable DB of type *sql.DB, which will hold a reference to the database connection pool
var DB *sql.DB

// initialize a connection to a SQLite database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("could not connect to database" + err.Error())
	}

	// Configure connection pool
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	createTables()
}

func createTables() {
	// create the users table if it doesn't already exist in the database
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("could not create user table: " + err.Error())
	}

	// create the events table if it doesn't already exist in the database
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("could not create events table: " + err.Error())
	}

}
