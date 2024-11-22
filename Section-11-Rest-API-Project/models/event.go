package models

import (
	"time"

	"example.com/building-a-rest-api/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	// add it to database
	insertQuery := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	// Prepare the SQL insert query
	stmt, err := db.DB.Prepare(insertQuery)

	if err != nil {
		return err
	}

	// Make sure to close the statement when the function exits
	defer stmt.Close()

	// Execute the insert query with the values from the Event struct
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	// Retrieve the ID of the newly inserted row
	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() []Event {
	return events
}
