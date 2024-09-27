package postgres

import (
	"database/sql"
	"fmt"
	"github.com/kksama1/DBCoursework/internal/model"
	_ "github.com/lib/pq"
	"io"
	"log"
	"os"
	"time"
)

type DatabaseDriver interface {
	SetUpDB()
	GetTables()
	CloseConnection() error
	InsertParticipant(participant model.Participant) (int, error)
	InsertVehicle(vehicle model.Vehicle) (int, error)
	InsertAccident(accident model.Accident) (int, error)
	InsertAccidentParticipant(accidentParticipant model.AccidentParticipant) (int, error)
}

var _ DatabaseDriver = (*PostgresDriver)(nil)

type PostgresDriver struct {
	Pool *sql.DB
}

// NewPostgresDriver is the constructor that  pointer to PostgresDriver instance.
func NewPostgresDriver(pool *sql.DB) *PostgresDriver {
	return &PostgresDriver{Pool: pool}
}

func CreateConnection(
	host string,
	port int,
	database string,
	username string,
	password string,
	sslmode string,
) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, username, password, database, sslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(15)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected to postgres!")
	}
	return db
}

func (p *PostgresDriver) CloseConnection() error {
	err := p.Pool.Close()
	if err != nil {
		return fmt.Errorf("error while closing conection: %v", err)
	}
	log.Println("connection closed")
	return nil
}

func (p *PostgresDriver) SetUpDB() {
	sqlFile, err := os.Open("/usr/local/src/db/sql/participants.sql")
	if err != nil {
		panic(err)
	}
	defer sqlFile.Close()

	sqlBytes, err := io.ReadAll(sqlFile)
	if err != nil {
		panic(err)
	}

	createTableQuery := string(sqlBytes)

	_, err = p.Pool.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

	sqlFile, err = os.Open("/usr/local/src/db/sql/drivers.sql")
	if err != nil {
		panic(err)
	}

	sqlBytes, err = io.ReadAll(sqlFile)
	if err != nil {
		panic(err)
	}

	createTableQuery = string(sqlBytes)

	_, err = p.Pool.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

	sqlFile, err = os.Open("/usr/local/src/db/sql/vehicles.sql")
	if err != nil {
		panic(err)
	}

	sqlBytes, err = io.ReadAll(sqlFile)
	if err != nil {
		panic(err)
	}

	createTableQuery = string(sqlBytes)

	_, err = p.Pool.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

	sqlFile, err = os.Open("/usr/local/src/db/sql/accidents.sql")
	if err != nil {
		panic(err)
	}

	sqlBytes, err = io.ReadAll(sqlFile)
	if err != nil {
		panic(err)
	}

	createTableQuery = string(sqlBytes)

	_, err = p.Pool.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

	sqlFile, err = os.Open("/usr/local/src/db/sql/accident_participants.sql")
	if err != nil {
		panic(err)
	}

	sqlBytes, err = io.ReadAll(sqlFile)
	if err != nil {
		panic(err)
	}

	createTableQuery = string(sqlBytes)

	_, err = p.Pool.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

}

func (p *PostgresDriver) GetTables() {
	log.Println("\tGetTables()")
	rows, err := p.Pool.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE'")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(tableName)
	}

	log.Println("\texit GetTables()")
}

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

func (p *PostgresDriver) GetAllAccidents() ([]model.Accident, error) {
	var accidents []model.Accident

	query := `SELECT accident_id, date, location, description FROM Accidents`

	rows, err := p.Pool.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	// Итерируем по результатам запроса
	for rows.Next() {
		var accident model.Accident
		err := rows.Scan(&accident.AccidentID, &accident.Date, &accident.Location, &accident.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		accidents = append(accidents, accident)
	}

	// Проверка на наличие ошибок после итерации
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return accidents, nil
}

type ByIDParticipant struct {
	ParticipantID int     `json:"participant_id" db:"participant_id"`
	FullName      string  `json:"full_name" db:"name"`
	IsDriver      bool    `json:"is_driver" db:"is_driver"`
	IsResponsible bool    `json:"is_responsible" db:"is_responsible"`
	VehicleID     *int    `json:"vehicle_id" db:"vehicle_id"`       // Указатель на VehicleID
	LicensePlate  *string `json:"license_plate" db:"license_plate"` // Указатель на LicensePlate
	Model         *string `json:"model" db:"model"`                 // Указатель на Model
}

// Функция для получения участников ДТП по ID
func (p *PostgresDriver) GetParticipantsByAccidentID(accidentID int) ([]ByIDParticipant, error) {
	var participants []ByIDParticipant

	query := `
	SELECT 
		p.participant_id, 
		p.name, 
		p.is_driver, 
		ap.is_responsible,
		v.vehicle_id,
		v.license_plate,
		v.model
	FROM Accident_Participants ap
	JOIN Participants p ON ap.participant_id = p.participant_id
	LEFT JOIN Vehicles v ON v.owner_id = p.participant_id
	WHERE ap.accident_id = $1
	`

	log.Printf("Executing query with accident_id: %d", accidentID)

	rows, err := p.Pool.Query(query, accidentID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var participant ByIDParticipant
		var vehicleID *int
		var licensePlate *string
		var model *string

		if err := rows.Scan(&participant.ParticipantID, &participant.FullName, &participant.IsDriver, &participant.IsResponsible, &vehicleID, &licensePlate, &model); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		// Устанавливаем информацию о транспортном средстве, если она доступна
		participant.VehicleID = vehicleID
		participant.LicensePlate = licensePlate
		participant.Model = model

		participants = append(participants, participant)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	log.Printf("Fetched %d participants", len(participants))
	return participants, nil
}

func (p *PostgresDriver) GetAccidentCount() (int, error) {
	var count int
	query := `
		SELECT 
			COUNT(*) AS total_accidents
		FROM Accidents;
	`

	err := p.Pool.QueryRow(query).Scan(&count)
	if err != nil {
		log.Println("Error executing query:", err)
		return 0, err
	}

	return count, nil
}

// ______________________________________________________________________________________________________________________

func translateDayOfWeek(day int) string {
	days := []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота", "Воскресенье"}
	if day < 0 || day > 6 {
		return "Неизвестный день"
	}
	return days[day]
}

// ______________________________________________________________________________________________________________________

func (p *PostgresDriver) GetAccidentReportByDayAndTime() ([]model.AccidentReport, error) {
	var accidents []model.AccidentReport

	// Запрос на получение ДТП по дням недели и времени суток
	query := `
		SELECT
			(EXTRACT(DOW FROM date) + 6) % 7 AS day_of_week, -- Преобразование дня недели, чтобы понедельник был первым
			COUNT(CASE WHEN EXTRACT(HOUR FROM date) BETWEEN 0 AND 5 THEN 1 END) AS night_count,
			COUNT(CASE WHEN EXTRACT(HOUR FROM date) BETWEEN 6 AND 11 THEN 1 END) AS morning_count,
			COUNT(CASE WHEN EXTRACT(HOUR FROM date) BETWEEN 12 AND 17 THEN 1 END) AS afternoon_count,
			COUNT(CASE WHEN EXTRACT(HOUR FROM date) BETWEEN 18 AND 23 THEN 1 END) AS evening_count
		FROM Accidents
		GROUP BY day_of_week
		ORDER BY day_of_week;
	`

	rows, err := p.Pool.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Считываем данные
	for rows.Next() {
		var accident model.AccidentReport
		var dayOfWeek int
		if err := rows.Scan(&dayOfWeek, &accident.NightCount, &accident.MorningCount, &accident.AfternoonCount, &accident.EveningCount); err != nil {
			return nil, err
		}

		// Преобразуем число дня недели в строку на русском
		accident.DayOfWeek = translateDayOfWeek(dayOfWeek)
		accidents = append(accidents, accident)
	}

	// Проверяем наличие ошибок после выполнения запроса
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accidents, nil
}
