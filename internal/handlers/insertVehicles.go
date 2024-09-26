package handlers

import (
	"github.com/kksama1/DBCoursework/internal/mocks"
	"log"
	"net/http"
)

func (s *Service) GeneratePlatese(w http.ResponseWriter, r *http.Request) {
	log.Println("\tGeneratePlatese()")

	for i := 0; i < 10; i++ {
		log.Println(mocks.GenerateLicensePlate())
	}

	log.Println("\texit GeneratePlatese()")

}
