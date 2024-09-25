package common

import (
	"database/sql"
	"github.com/kksama1/DBCoursework/internal/db/postgres"
)

type ServiceHandler interface {
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
