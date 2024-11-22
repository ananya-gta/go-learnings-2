package models

import (
	"time"

	"example.com/building-a-rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time 
	UserID      int
}

var events = []Event{}

func (e Event) Save() error{
	// add it to database
	insertQuery := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(insertQuery)

	if err != nil {
		return err
	}

	defer stmt.Close()

	// execute the insert query
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err
}

func GetAllEvents() []Event {
	return events
}
