package handlers

import (
	"fmt"
	"github.com/kksama1/DBCoursework/internal/mocks"
	"github.com/kksama1/DBCoursework/internal/model"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func (s *Service) GenerateTest(w http.ResponseWriter, r *http.Request) {
	//rand.Seed(time.Now().UnixNano())
	//
	//// Генерируем случайные данные
	//participants := make([]model.Participant, 10)
	//vehicles := make([]model.Vehicle, 10)
	//drivers := make([]model.Driver, 10)
	//accidents := make([]model.Accident, 5)
	//accidentParticipants := make([]model.AccidentParticipant, 10)
	//
	//// Заполняем участников
	//for i := 0; i < 10; i++ {
	//	participants[i] = mocks.CreateRandomParticipant()
	//	fmt.Printf("Participant: %+v\n", participants[i])
	//}
	//
	//// Заполняем транспортные средства
	//for i := 0; i < 10; i++ {
	//	vehicles[i] = mocks.CreateRandomVehicle(participants[i].ParticipantID)
	//	fmt.Printf("Vehicle: %+v\n", vehicles[i])
	//}
	//
	//// Заполняем водителей
	//for i := 0; i < 10; i++ {
	//	drivers[i] = mocks.CreateRandomDriver(participants[i].ParticipantID)
	//	fmt.Printf("Driver: %+v\n", drivers[i])
	//}
	//
	//// Заполняем ДТП
	//for i := 0; i < 5; i++ {
	//	accidents[i] = mocks.CreateRandomAccident()
	//	fmt.Printf("Accident: %+v\n", accidents[i])
	//}
	//
	//// Заполняем участников ДТП
	//for i := 0; i < 5; i++ {
	//	accident := mocks.CreateRandomAccident()
	//	accidents[i] = accident
	//	fmt.Printf("Accident: %+v\n", accident)
	//
	//	// Генерируем участников для каждой аварии
	//	numParticipants := rand.Intn(3) + 2 // Минимум 2, максимум 4 участника
	//
	//	// Переменные для учета водителей
	//	var responsibleParticipantID int
	//	var isResponsibleFound bool
	//
	//	for j := 0; j < numParticipants; j++ {
	//		// Определяем, является ли участник пешеходом (каждый пятый)
	//		isDriver := j%5 != 0 // Каждый 5-й участник будет пешеходом
	//		var vehicleID *int   // Указатель на int для vehicleID
	//
	//		// Случайно выбираем участника
	//		participant := participants[rand.Intn(len(participants))]
	//
	//		// Если это водитель, получаем ID автомобиля
	//		if isDriver {
	//			randomVehicle := vehicles[rand.Intn(len(vehicles))] // Случайный автомобиль
	//			vehicleID = &randomVehicle.VehicleID
	//
	//			// Определяем, будет ли этот водитель ответственным
	//			if !isResponsibleFound {
	//				isResponsibleFound = true
	//				responsibleParticipantID = participant.ParticipantID
	//			}
	//		}
	//
	//		accidentParticipant := mocks.CreateRandomAccidentParticipant(
	//			accident.AccidentID,
	//			participant.ParticipantID,
	//			isDriver,
	//			vehicleID,
	//			// Устанавливаем isResponsible в true, если это ответственный водитель
	//			participant.ParticipantID == responsibleParticipantID,
	//		)
	//		accidentParticipants = append(accidentParticipants, accidentParticipant)
	//		fmt.Printf("AccidentParticipant: %+v\n", accidentParticipant)
	//	}
	//}

	//rand.Seed(time.Now().UnixNano())
	//// Генерация участников
	//const numParticipants = 10
	//participants := make([]model.Participant, numParticipants)
	//vehicles := make([]model.Vehicle, 0)
	//var accidentID int
	//
	//// Создание участников
	//for i := 0; i < numParticipants; i++ {
	//	participant := mocks.CreateRandomParticipant()
	//	participantID, err := s.DB.InsertParticipant(participant)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	participants[i] = participant
	//
	//	// Если участник является водителем, создаем транспортное средство
	//	if participant.IsDriver {
	//		vehicle := mocks.CreateRandomVehicle(participantID)
	//		vehicleID, err := s.DB.InsertVehicle(vehicle)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		vehicles = append(vehicles, vehicle)
	//		fmt.Printf("Inserted Vehicle with ID: %d\n", vehicleID)
	//	}
	//}
	//
	//// Создание ДТП
	//accident := mocks.CreateRandomAccident()
	//accidentID, err := s.DB.InsertAccident(accident)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Inserted Accident with ID: %d\n", accidentID)
	//
	//// Создание участников ДТП
	//responsibleParticipantID := -1 // ID ответственного участника
	//
	//for _, participant := range participants {
	//	// Устанавливаем responsibleParticipantID для водителей
	//	if participant.IsDriver && responsibleParticipantID == -1 {
	//		responsibleParticipantID = participant.ParticipantID
	//	}
	//
	//	var vehicleID *int
	//	if participant.IsDriver {
	//		for _, vehicle := range vehicles {
	//			if vehicle.OwnerID == participant.ParticipantID {
	//				vehicleID = &vehicle.VehicleID
	//				break
	//			}
	//		}
	//	}
	//
	//	// Устанавливаем isResponsible для ответственного участника
	//	isResponsible := participant.ParticipantID == responsibleParticipantID
	//
	//	// Проверяем, что participant_id существует перед вставкой
	//	fmt.Printf("Inserting AccidentParticipant for participant ID: %d\n", participant.ParticipantID)
	//	accidentParticipant := mocks.CreateRandomAccidentParticipant(accidentID, participant.ParticipantID, participant.IsDriver, vehicleID, isResponsible)
	//	accidentParticipantID, err := s.DB.InsertAccidentParticipant(db, accidentParticipant)
	//	if err != nil {
	//		log.Fatalf("Failed to insert AccidentParticipant: %v", err)
	//	}
	//	fmt.Printf("Inserted AccidentParticipant with ID: %d\n", accidentParticipantID)
	//}

	rand.Seed(time.Now().UnixNano())

	const numParticipants = 10
	participants := make([]model.Participant, 0, numParticipants) // Изменяем размер массива
	vehicles := make([]model.Vehicle, 0)
	var accidentID int
	var responsibleParticipantID int // ID ответственного участника

	// Создание участников
	for i := 0; i < numParticipants; i++ {
		participant := mocks.CreateRandomParticipant()
		participantID, err := s.DB.InsertParticipant(participant)
		if err != nil {
			log.Fatal(err)
		}
		participant.ParticipantID = participantID        // Обновляем ID у созданного участника
		participants = append(participants, participant) // добавляем участника в массив
		fmt.Printf("Inserted Participant with ID: %d\n", participantID)

		// Если участник является водителем, создаем транспортное средство
		if participant.IsDriver {
			vehicle := mocks.CreateRandomVehicle(participantID)
			vehicleID, err := s.DB.InsertVehicle(vehicle)
			if err != nil {
				log.Fatal(err)
			}
			vehicle.VehicleID = vehicleID
			vehicles = append(vehicles, vehicle)
			fmt.Printf("Inserted Vehicle with ID: %d\n", vehicleID)

			// Создаем запись о водителе
			driver := mocks.CreateRandomDriver(participantID)
			_, err = s.DB.InsertDriver(driver)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Inserted Driver with ID: %d\n", participantID)
		}
	}

	// Создание ДТП
	accident := mocks.CreateRandomAccident()
	accidentID, err := s.DB.InsertAccident(accident)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted Accident with ID: %d\n", accidentID)

	// Создание участников ДТП
	fmt.Println("Participants IDs:")
	for _, participant := range participants {
		fmt.Printf("Participant ID: %d\n", participant.ParticipantID) // Проверяем все ID участников
	}

	// Устанавливаем ответственного участника
	for _, participant := range participants {
		if participant.IsDriver && responsibleParticipantID == 0 { // 0 по умолчанию, а не -1
			responsibleParticipantID = participant.ParticipantID
		}
	}

	// Вставка участников ДТП
	for _, participant := range participants {
		var vehicleID *int // По умолчанию это nil для пешеходов (участников без ТС)

		if participant.IsDriver {
			// Ищем транспортное средство, если участник — водитель
			for _, vehicle := range vehicles {
				if vehicle.OwnerID == participant.ParticipantID {
					vehicleID = &vehicle.VehicleID // Устанавливаем vehicleID, если участник является водителем
					break
				}
			}

			// Если водитель, но не найдено транспортное средство — это ошибка
			if vehicleID == nil {
				log.Printf("Error: No vehicle found for driver with participant ID: %d", participant.ParticipantID)
				continue // Пропускаем участника, если он водитель, но у него нет ТС
			}
		}

		// Устанавливаем ответственность (isResponsible) для ответственного участника
		isResponsible := participant.ParticipantID == responsibleParticipantID

		// Вставляем участника ДТП
		var accidentParticipant = model.AccidentParticipant{}
		fmt.Printf("Inserting AccidentParticipant for participant ID: %d\n", participant.ParticipantID)
		if participant.IsDriver && vehicleID != nil {
			accidentParticipant = mocks.CreateRandomAccidentParticipant(accidentID, participant.ParticipantID, participant.IsDriver, vehicleID, isResponsible)
		} else {
			accidentParticipant = mocks.CreateRandomAccidentParticipant(accidentID, participant.ParticipantID, participant.IsDriver, nil, isResponsible)
		}
		accidentParticipantID, err := s.DB.InsertAccidentParticipant(accidentParticipant)
		if err != nil {
			log.Fatalf("Failed to insert AccidentParticipant: %v", err)
		}

		fmt.Printf("Inserted AccidentParticipant with ID: %d\n", accidentParticipantID)
	}

}
