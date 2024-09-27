package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// GetAllAccidents обрабатывает запрос на получение всех ДТП и возвращает HTML-страницу с таблицей.
func (s *Service) GetAllAccidents(w http.ResponseWriter, r *http.Request) {
	accidents, err := s.DB.GetAllAccidents()
	if err != nil {
		log.Println(err)
		http.Error(w, "не удалось получить данные о ДТП", http.StatusInternalServerError)
		return
	}

	// Формирование HTML-страницы с таблицей
	tmpl := `
	<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Все ДТП</title>
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
			.clickable-row {
				cursor: pointer;
			}
		</style>
	</head>
	<body>
		<h1>Все ДТП</h1>
		<table>
			<tr>
				<th>ID ДТП</th>
				<th>Дата</th>
				<th>Место</th>
				<th>Описание</th>
			</tr>
			{{range .}}
			<tr class="clickable-row" onclick="window.location='http://localhost:8080/getParticipantsByAccidentIDHandler?accident_id={{.AccidentID}}'">
				<td>{{.AccidentID}}</td>
				<td>{{.Date.Format "2006-01-02 15:04:05"}}</td>
				<td>{{.Location}}</td>
				<td>{{.Description}}</td>
			</tr>
			{{end}}
		</table>
	</body>
	</html>
	`

	// Создание и выполнение шаблона
	t, err := template.New("accidents").Parse(tmpl)
	if err != nil {
		log.Printf("Ошибка при разборе шаблона: %v", err)
		http.Error(w, "не удалось сгенерировать ответ", http.StatusInternalServerError)
		return
	}

	// Установка заголовков и возврат результата
	w.Header().Set("Content-Type", "text/html")
	if err := t.Execute(w, accidents); err != nil {
		log.Printf("Ошибка при выполнении шаблона: %v", err)
		http.Error(w, "не удалось сгенерировать ответ", http.StatusInternalServerError)
		return
	}
}
