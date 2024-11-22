package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// declare a global variable DB of type *sql.DB, which will hold a reference to the database connection pool
var DB *sql.DB

//  initialize a connection to a SQLite database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}

