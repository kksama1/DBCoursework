package mocks

import (
	"fmt"
	"github.com/kksama1/DBCoursework/internal/model"
	"math/rand"
	"time"
)

// Хардкод имен и фамилий
var firstNames = []string{"Иван", "Петр", "Сергей", "Алексей", "Дмитрий"}
var lastNames = []string{"Иванов", "Петров", "Сидоров", "Кузнецов", "Смирнов", "Попов", "Лебедев", "Морозов", "Васильев", "Зайцев"}

// Функция для генерации случайного имени
func randomFullName() string {
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	return fmt.Sprintf("%s %s", firstName, lastName)
}

// Функция для генерации случайной строки
func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// Функция для генерации случайного участника
func CreateRandomParticipant() model.Participant {
	name := randomFullName()
	dateOfBirth := time.Now().AddDate(-rand.Intn(30), 0, 0) // Возраст до 30 лет
	isDriver := rand.Intn(4) != 1                           // Случайно назначаем, является ли водителем

	return model.Participant{
		ParticipantID: rand.Intn(1000), // Заменить на логику генерации уникальных ID
		FullName:      name,
		DateOfBirth:   dateOfBirth,
		LicenseNumber: randomString(5) + "12345", // Случайный номер прав
		IsDriver:      isDriver,
	}
}

// Функция для генерации случайного транспортного средства
// Функция для генерации случайного транспортного средства
func CreateRandomVehicle(ownerID int) model.Vehicle {
	//licensePlate := randomString(3) + "-" + randomString(3) + "-" + randomString(2)
	vehicleModel := randomString(7) // Изменено имя переменной для предотвращения конфликта
	year := 2000 + rand.Intn(23)    // Год от 2000 до 2022

	return model.Vehicle{
		VehicleID:    rand.Intn(1000), // Заменить на логику генерации уникальных ID
		LicensePlate: GenerateLicensePlate(),
		Model:        vehicleModel, // Используем новое имя переменной
		Year:         year,
		OwnerID:      ownerID,
	}
}

// Функция для генерации случайного водителя
func CreateRandomDriver(participantID int) model.Driver {
	licenseNumber := randomString(5) + "12345"
	experienceYears := rand.Intn(10) // Стаж от 0 до 9 лет

	return model.Driver{
		DriverID:        participantID,
		LicenseNumber:   licenseNumber,
		ExperienceYears: experienceYears,
	}
}

// Функция для генерации случайного ДТП
func CreateRandomAccident() model.Accident {
	date := time.Now()
	location := randomString(15)
	description := randomString(50)

	return model.Accident{
		AccidentID:  rand.Intn(1000), // Заменить на логику генерации уникальных ID
		Date:        date,
		Location:    location,
		Description: description,
	}
}

// Функция для генерации случайного участника ДТП
func CreateRandomAccidentParticipant(accidentID, participantID int, isDriver bool, vehicleID *int, isResponsible bool) model.AccidentParticipant {
	role := "водитель"

	// Если это пешеход, устанавливаем role и isResponsible
	if !isDriver {
		role = "пешеход"
		vehicleID = nil // Установим vehicleID в nil, если это пешеход
	}

	return model.AccidentParticipant{
		AccidentParticipantID: rand.Intn(1000), // Заменить на логику генерации уникальных ID
		ParticipantID:         participantID,
		AccidentID:            accidentID,
		VehicleID:             vehicleID,
		IsResponsible:         isResponsible,
		Role:                  role,
	}
}
