package models

import (
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

func (e Event) Save()  {
	events = append(events, e)
}

func GetAllEvent() []Event {
	return events
}

func AddNewEvent(e Event) {

}
