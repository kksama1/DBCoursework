package main

import (
	"github.com/kksama1/DBCoursework/internal/common"
	"github.com/kksama1/DBCoursework/internal/config"
	"github.com/kksama1/DBCoursework/internal/db/postgres"
	"log"
	"net/http"
)

func main() {
	log.Println("server started!")

	cfg, err := config.LoadConfig[config.DatabaseConfig]()
	if err != nil {
		log.Fatal(err)
	}
	pool := postgres.CreateConnection(cfg.Host, cfg.Port, cfg.Database, cfg.Username, cfg.Password, cfg.Sslmode)
	service := common.NewService(pool)
	defer func() {
		err = service.DB.CloseConnection()
		if err != nil {
			log.Println(err)
		}
	}()

	service.DB.SetUpDB()
	service.DB.GetTables()

	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
