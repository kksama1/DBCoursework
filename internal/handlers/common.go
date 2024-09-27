package handlers

import (
	"database/sql"
	"github.com/kksama1/DBCoursework/internal/db/postgres"
	"net/http"
)

type ServiceHandler interface {
	GeneratePlatese(w http.ResponseWriter, r *http.Request)
	GenerateTest(w http.ResponseWriter, r *http.Request)
	GetAllAccidents(w http.ResponseWriter, r *http.Request)
	GetParticipantsByAccidentIDHandler(w http.ResponseWriter, r *http.Request)
	TotalAccidentHandler(w http.ResponseWriter, r *http.Request)
	GetAccidentReportByDayAndTimeHandler(w http.ResponseWriter, r *http.Request)
}

var _ ServiceHandler = (*Service)(nil)

type Service struct {
	DB *postgres.PostgresDriver
}

func NewService(pool *sql.DB) *Service {
	return &Service{
		DB: postgres.NewPostgresDriver(pool),
	}
}
