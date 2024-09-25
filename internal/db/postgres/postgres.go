package postgres

import (
	"database/sql"
	"fmt"
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
	rows, err := p.Pool.Query("SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE'")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	// Чтение результатов запроса
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
