package models

import (
	"sum/gin-api/db"
	"time"
)

type Event struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime"`
	UserId      int       `json:"user_id"`
}

var events = []Event{}

func (e Event) Save() error {
	//simulating DB ops
	query := `
	INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	e.Id = id
	events = append(events, e)
	return nil
}

func GetAllEvent() []Event {
	return events
}

func AddNewEvent(e Event) {

}
