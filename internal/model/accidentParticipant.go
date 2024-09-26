package model

type AccidentParticipant struct {
	AccidentParticipantID int    `json:"accident_participant_id" db:"accident_participant_id"`
	ParticipantID         int    `json:"participant_id" db:"participant_id"` // Ссылка на Participants
	AccidentID            int    `json:"accident_id" db:"accident_id"`       // Ссылка на Accidents
	VehicleID             *int   `json:"vehicle_id" db:"vehicle_id"`         // Может быть nil для пешеходов
	IsResponsible         bool   `json:"is_responsible" db:"is_responsible"` // Ответственный за ДТП
	Role                  string `json:"role" db:"role"`                     // Роль участника: 'водитель' или 'пешеход'
}
