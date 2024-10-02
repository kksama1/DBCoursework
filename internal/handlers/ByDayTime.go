package handlers

import (
	"html/template"
	"net/http"
)

// GetAccidentReportByDayAndTimeHandler обрабатывает запрос на получение отчета по ДТП по дням недели и времени суток.
func (s *Service) GetAccidentReportByDayAndTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем данные из базы
	accidents, err := s.DB.GetAccidentReportByDayAndTime()
	if err != nil {
		http.Error(w, "Failed to fetch accidents", http.StatusInternalServerError)
		return
	}

	// Определяем HTML-шаблон с навигационной панелью
	tmpl := `
	<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>ДТП по дням и времени</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				margin: 20px;
			}
			h1 {
				color: #333;
			}
			/* Стили для таблицы */
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
			/* Стили для навигационной панели */
			.navbar {
				overflow: hidden;
				background-color: #f2f2f2;
				margin-bottom: 20px;
				border-bottom: 1px solid #ddd;
			}
			.navbar a {
				float: left;
				display: block;
				color: black;
				text-align: center;
				padding: 14px 16px;
				text-decoration: none;
				border-right: 1px solid #ddd;
			}
			.navbar a:last-child {
				border-right: none;
			}
			.navbar a:hover {
				background-color: #f5f5f5;
				color: black;
			}
		</style>
	</head>
	<body>
		<!-- Навигационная панель -->
		<div class="navbar">
			<a href="/getAllAccidents">Все ДТП</a>
			<a href="/totalAccidents">Общее количество ДТП</a>
			<a href="/byTime">ДТП по дням недели и времени суток</a>
		</div>

		<h1>ДТП по дням недели и времени суток</h1>
		<table>
			<tr>
				<th>День недели</th>
				<th>Ночь</th>
				<th>Утро</th>
				<th>День</th>
				<th>Вечер</th>
			</tr>
			{{range .}}
			<tr>
				<td>{{.DayOfWeek}}</td>
				<td>{{.NightCount}}</td>
				<td>{{.MorningCount}}</td>
				<td>{{.AfternoonCount}}</td>
				<td>{{.EveningCount}}</td>
			</tr>
			{{end}}
		</table>
	</body>
	</html>
	`

	// Создаем новый шаблон и парсим его
	t, err := template.New("accidents").Parse(tmpl)
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	// Генерируем HTML и отправляем его в ответ
	w.Header().Set("Content-Type", "text/html")
	if err := t.Execute(w, accidents); err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}
