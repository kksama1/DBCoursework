package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// GetParticipantsByAccidentIDHandler обрабатывает запрос на получение участников ДТП по идентификатору ДТП и возвращает HTML-страницу с таблицей.
func (s *Service) GetParticipantsByAccidentIDHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение параметра accident_id из URL
	accidentIDParam := r.URL.Query().Get("accident_id")
	if accidentIDParam == "" {
		http.Error(w, "accident_id is required", http.StatusBadRequest)
		return
	}

	// Преобразование accident_id в int
	accidentID, err := strconv.Atoi(accidentIDParam)
	if err != nil {
		http.Error(w, "invalid accident_id", http.StatusBadRequest)
		return
	}

	// Извлечение участников по идентификатору ДТП
	participants, err := s.DB.GetParticipantsByAccidentID(accidentID)
	if err != nil {
		log.Printf("Ошибка при получении участников ДТП для accident ID %d: %v", accidentID, err)
		http.Error(w, "не удалось получить участников", http.StatusInternalServerError)
		return
	}

	// Проверка, найдены ли участники
	if len(participants) == 0 {
		http.Error(w, "участники не найдены для данного accident_id", http.StatusNotFound)
		return
	}

	// Формирование HTML-страницы с таблицей
	tmpl := `
	<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Участники ДТП</title>
		<style>
			table {
				width: 100%;
				border-collapse: collapse;
			}
			th, td {
				border: 1px solid #ddd;
				padding: 8px;
				text-align: center;
			}
			th {
				background-color: #f2f2f2;
			}
			tr:hover {
				background-color: #f5f5f5;
			}
			.responsible {
				background-color: #ffcccc; /* Светло-красный цвет для виновников */
			}
		</style>
	</head>
	<body>
		<h1>Участники ДТП</h1>
		<table>
			<tr>
				<th>ID Участника</th>
				<th>ФИО</th>
				<th>Водитель</th>
				<th>Виновник ДТП</th>
				<th>Модель автомобиля</th>
				<th>Номер автомобиля</th>
			</tr>
			{{range .}}
			<tr class="{{if .IsResponsible}}responsible{{end}}">
				<td>{{.ParticipantID}}</td>
				<td>{{.FullName}}</td>
				<td>{{if .IsDriver}}Да{{else}}Нет{{end}}</td>
				<td>{{if .IsResponsible}}Да{{else}}Нет{{end}}</td>
				<td>{{if .Model}}{{.Model}}{{else}}Нет{{end}}</td>
				<td>{{if .LicensePlate}}{{.LicensePlate}}{{else}}Нет{{end}}</td>
			</tr>
			{{end}}
		</table>
	</body>
	</html>
	`

	// Создание и выполнение шаблона
	t, err := template.New("participants").Parse(tmpl)
	if err != nil {
		log.Printf("Ошибка при разборе шаблона: %v", err)
		http.Error(w, "не удалось сгенерировать ответ", http.StatusInternalServerError)
		return
	}

	// Установка заголовков и возврат результата
	w.Header().Set("Content-Type", "text/html")
	if err := t.Execute(w, participants); err != nil {
		log.Printf("Ошибка при выполнении шаблона: %v", err)
		http.Error(w, "не удалось сгенерировать ответ", http.StatusInternalServerError)
		return
	}
}
