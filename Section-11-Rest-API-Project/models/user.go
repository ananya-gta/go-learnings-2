package models

import (
	"errors"

	"example.com/building-a-rest-api/db"
	"example.com/building-a-rest-api/utils"
)

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

	// hash the password before saving it into the database
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	// get the last insert id
	lastID, err := result.LastInsertId()
	// update the user id
	u.ID = lastID
	
	return err
}

func (u User) ValidateUser() error {
	// retrieve user from the database
	query := "SELECT email, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return err
	}

	isValidPassword := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if  !isValidPassword {
		return errors.New("invalid password")
	}

	return nil
}
