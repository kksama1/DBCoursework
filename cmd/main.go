package main

import (
	"github.com/kksama1/DBCoursework/internal/config"
	"github.com/kksama1/DBCoursework/internal/db/postgres"
	"github.com/kksama1/DBCoursework/internal/handlers"
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
	service := handlers.NewService(pool)
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

	router.HandleFunc("/generateNums", service.GeneratePlatese)
	router.HandleFunc("/test", service.GenerateTest)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
