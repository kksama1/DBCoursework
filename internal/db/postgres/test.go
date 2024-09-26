package postgres

import (
	"database/sql"
	"github.com/kksama1/DBCoursework/internal/model"
	"log"
)

// Функция для вставки участника
func (p *PostgresDriver) InsertParticipant(participant model.Participant) (int, error) {
	var participantID int
	query := `INSERT INTO Participants (name, date_of_birth, license_number, is_driver) 
              VALUES ($1, $2, $3, $4) RETURNING participant_id`
	err := p.Pool.QueryRow(query, participant.FullName, participant.DateOfBirth, participant.LicenseNumber, participant.IsDriver).Scan(&participantID)
	if err != nil {
		return 0, err
	}
	return participantID, nil
}

func (p *PostgresDriver) InsertVehicle(vehicle model.Vehicle) (int, error) {
	var vehicleID int
	query := `INSERT INTO Vehicles (license_plate, model, year, owner_id) 
              VALUES ($1, $2, $3, $4) RETURNING vehicle_id`
	err := p.Pool.QueryRow(query, vehicle.LicensePlate, vehicle.Model, vehicle.Year, vehicle.OwnerID).Scan(&vehicleID)
	if err != nil {
		return 0, err
	}
	return vehicleID, nil
}

func (p *PostgresDriver) InsertAccident(accident model.Accident) (int, error) {
	var accidentID int
	query := `INSERT INTO Accidents (date, location, description) 
              VALUES ($1, $2, $3) RETURNING accident_id`
	err := p.Pool.QueryRow(query, accident.Date, accident.Location, accident.Description).Scan(&accidentID)
	if err != nil {
		return 0, err
	}
	return accidentID, nil
}

func (p *PostgresDriver) InsertAccidentParticipant(accidentParticipant model.AccidentParticipant) (int, error) {
	var accidentParticipantID int

	query := `INSERT INTO Accident_Participants (participant_id, accident_id, vehicle_id, is_responsible, role) 
              VALUES ($1, $2, $3, $4, $5) RETURNING accident_participant_id`

	// Обрабатываем NULL для vehicle_id
	var vehicleID sql.NullInt32
	if accidentParticipant.VehicleID != nil {
		vehicleID = sql.NullInt32{Int32: int32(*accidentParticipant.VehicleID), Valid: true}
	} else {
		vehicleID = sql.NullInt32{Int32: 0, Valid: false} // Используем NULL для пешеходов
	}

	err := p.Pool.QueryRow(query,
		accidentParticipant.ParticipantID,
		accidentParticipant.AccidentID,
		vehicleID, // Здесь теперь безопасно
		accidentParticipant.IsResponsible,
		accidentParticipant.Role).Scan(&accidentParticipantID)

	// Логирование для диагностики
	log.Println("Accident_Participant", accidentParticipant)
	if accidentParticipant.VehicleID != nil {
		log.Println("Vehicle ID:", *accidentParticipant.VehicleID)
	} else {
		log.Println("No vehicle for this participant (pedestrian)")
	}

	if err != nil {
		return 0, err
	}

	return accidentParticipantID, nil
}

func (p *PostgresDriver) InsertDriver(driver model.Driver) (int, error) {
	var driverID int
	query := `
		INSERT INTO Drivers (driver_id, license_number, experience_years) 
		VALUES ($1, $2, $3) RETURNING driver_id`

	err := p.Pool.QueryRow(query, driver.DriverID, driver.LicenseNumber, driver.ExperienceYears).Scan(&driverID)
	if err != nil {
		return 0, err
	}
	return driverID, nil
}
