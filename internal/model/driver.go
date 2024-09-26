package model

type Driver struct {
	DriverID        int    `json:"driver_id" db:"driver_id"` // Ссылка на Participants
	LicenseNumber   string `json:"license_number" db:"license_number"`
	ExperienceYears int    `json:"experience_years" db:"experience_years"`
}
