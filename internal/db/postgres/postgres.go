package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type DatabaseDriver interface {
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
