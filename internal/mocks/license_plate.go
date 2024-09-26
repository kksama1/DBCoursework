package mocks

import (
	"fmt"
	"math/rand"
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
