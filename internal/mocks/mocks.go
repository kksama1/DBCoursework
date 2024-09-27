package mocks

import (
	"fmt"
	"github.com/kksama1/DBCoursework/internal/model"
	"math/rand"
	"time"
)

var letters = []rune{'А', 'В', 'Е', 'К', 'М', 'Н', 'О', 'Р', 'С', 'Т', 'У', 'Х'}

// Генерация случайного автомобильного номера
func GenerateLicensePlate() string {
	// Генерация букв
	firstLetter := letters[rand.Intn(len(letters))]
	secondLetter := letters[rand.Intn(len(letters))]
	thirdLetter := letters[rand.Intn(len(letters))]

	// Генерация трёх цифр
	number := rand.Intn(900) + 100 // случайное число от 100 до 999

	// Генерация кода региона
	region := rand.Intn(999) + 1 // случайное число от 1 до 999

	// Формирование итогового номера
	return fmt.Sprintf("%c%d%c%c %d", firstLetter, number, secondLetter, thirdLetter, region)
}

// Хардкод имен и фамилий
var firstNames = []string{"Иван", "Петр", "Сергей", "Алексей", "Дмитрий"}
var lastNames = []string{"Иванов", "Петров", "Сидоров", "Кузнецов", "Смирнов", "Попов", "Лебедев", "Морозов", "Васильев", "Зайцев"}

var streets = []string{
	"Улица Ленина",
	"Улица Пушкина",
	"Проспект Мира",
	"Улица Гагарина",
	"Улица Солнечная",
	"Улица Мира",
	"Улица Рябиновая",
	"Улица Цветочная",
	"Проспект Победы",
	"Улица Московская",
}

var violations = []string{
	"превышение скорости",
	"непредоставление преимущества в движении",
	"проезд на красный свет",
	"вождение в состоянии алкогольного опьянения",
	"неиспользование ремня безопасности",
	"выезд на встречную полосу",
	"нарушение правил парковки",
	"отсутствие включенных фар в темное время суток",
	"разговор по телефону за рулем",
	"переполнение транспортного средства",
}

// Примеры марок автомобилей
var vehicleBrands = []string{
	"Toyota",
	"Ford",
	"Volkswagen",
	"Chevrolet",
	"Honda",
	"Nissan",
	"Hyundai",
	"Subaru",
	"BMW",
	"Mercedes-Benz",
}

// Примеры моделей автомобилей
var vehicleModels = map[string][]string{
	"Toyota":        {"Corolla", "Camry", "Prius", "RAV4"},
	"Ford":          {"Focus", "Mustang", "Explorer", "Fiesta"},
	"Volkswagen":    {"Golf", "Jetta", "Passat", "Tiguan"},
	"Chevrolet":     {"Malibu", "Impala", "Cruze", "Equinox"},
	"Honda":         {"Civic", "Accord", "CR-V", "Fit"},
	"Nissan":        {"Altima", "Sentra", "Rogue", "Murano"},
	"Hyundai":       {"Elantra", "Sonata", "Tucson", "Santa Fe"},
	"Subaru":        {"Impreza", "Legacy", "Outback", "Forester"},
	"BMW":           {"3 Series", "5 Series", "X3", "X5"},
	"Mercedes-Benz": {"C-Class", "E-Class", "GLC", "S-Class"},
}

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
func CreateRandomVehicle(ownerID int) model.Vehicle {
	// Выбор случайной марки автомобиля
	brand := vehicleBrands[rand.Intn(len(vehicleBrands))]
	// Выбор случайной модели из выбранной марки
	modelName := vehicleModels[brand][rand.Intn(len(vehicleModels[brand]))]
	year := 2000 + rand.Intn(23) // Год от 2000 до 2022

	// Объединяем марку и модель в одну строку
	fullModel := fmt.Sprintf("%s %s", brand, modelName)

	return model.Vehicle{
		VehicleID:    rand.Intn(1000), // Заменить на логику генерации уникальных ID
		LicensePlate: GenerateLicensePlate(),
		Model:        fullModel, // Используем объединенную строку марки и модели
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
	// Текущее время
	now := time.Now()

	// 15 лет в днях (15 лет * 365 дней, не учитывая високосные года для простоты)
	daysIn15Years := 15 * 365

	// Генерируем случайное количество дней назад (в пределах 15 лет)
	randomDays := rand.Intn(daysIn15Years)

	// Вычитаем это количество дней из текущего времени
	randomDate := now.AddDate(0, 0, -randomDays)

	// Генерируем случайное время в пределах суток для этой даты
	randomHour := rand.Intn(24)   // Часы от 0 до 23
	randomMinute := rand.Intn(60) // Минуты от 0 до 59
	randomSecond := rand.Intn(60) // Секунды от 0 до 59

	// Добавляем случайное время к случайной дате
	randomDate = time.Date(randomDate.Year(), randomDate.Month(), randomDate.Day(), randomHour, randomMinute, randomSecond, 0, randomDate.Location())

	// Генерация случайного местоположения и описания
	location := streets[rand.Intn(len(streets))]
	description := violations[rand.Intn(len(violations))]

	// Возвращаем случайный объект ДТП
	return model.Accident{
		AccidentID:  rand.Intn(1000), // Можно заменить на более продвинутую логику генерации уникальных ID
		Date:        randomDate,
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
