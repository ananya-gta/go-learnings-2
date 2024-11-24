package models

import "example.com/building-a-rest-api/db"

type User struct {
	ID       int64
	Email    string `bindng:"required"`
	Password string `bindng:"required"`
}

func (u User) Save() error {
	// add user to the database

	query := "INSERT INTO users (email, password) VALUES (?, ?)"

	// prepare the insert query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Make sure to close the statement when the function exits
	defer stmt.Close()

	result, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	// get the last insert id
	lastID, err := result.LastInsertId()
	// update the user id
	u.ID = lastID
	
	return err
}
