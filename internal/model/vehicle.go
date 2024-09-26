package model

type Vehicle struct {
	VehicleID    int    `json:"vehicle_id" db:"vehicle_id"`
	LicensePlate string `json:"license_plate" db:"license_plate"`
	Model        string `json:"model" db:"model"`
	Year         int    `json:"year" db:"year"`
	OwnerID      int    `json:"owner_id" db:"owner_id"` // Ссылка на Participants
}
