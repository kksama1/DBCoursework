package model

import "time"

type Accident struct {
	AccidentID  int       `json:"accident_id" db:"accident_id"`
	Date        time.Time `json:"date" db:"date"`
	Location    string    `json:"location" db:"location"`
	Description string    `json:"description" db:"description"`
}
