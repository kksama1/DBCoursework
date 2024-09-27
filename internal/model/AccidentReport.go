package model

// AccidentReport - структура для хранения отчёта по ДТП
type AccidentReport struct {
	DayOfWeek      string
	NightCount     int
	MorningCount   int
	AfternoonCount int
	EveningCount   int
}
