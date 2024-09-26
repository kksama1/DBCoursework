package model

import "time"

type Participant struct {
	ParticipantID int       `json:"participant_id" db:"participant_id"`
	FullName      string    `json:"full_name" db:"name"`
	DateOfBirth   time.Time `json:"date_of_birth" db:"date_of_birth"`
	LicenseNumber string    `json:"license_number" db:"license_number"` // Может быть пустым
	IsDriver      bool      `json:"is_driver" db:"is_driver"`
}
