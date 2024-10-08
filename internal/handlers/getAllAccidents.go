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

	// Формирование HTML-страницы с таблицей и навигационной панелью
	tmpl := `
	<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Все ДТП</title>
		<style>
			/* Основные стили для таблицы */
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
				background-color: #f2f2f2; /* тот же цвет для заголовка таблицы */
			}
			tr:hover {
				background-color: #f5f5f5;
			}
			.clickable-row {
				cursor: pointer;
			}

			/* Стили для навигационной панели, схожие с заголовком таблицы */
			.navbar {
				overflow: hidden;
				background-color: #f2f2f2; /* тот же цвет, что и у заголовков таблицы */
				margin-bottom: 20px;
				border-bottom: 1px solid #ddd;
			}
			.navbar a {
				float: left;
				display: block;
				color: black; /* Текст ссылки черного цвета */
				text-align: center;
				padding: 14px 16px;
				text-decoration: none;
				border-right: 1px solid #ddd; /* Добавляем границу, как у заголовков таблицы */
			}
			.navbar a:last-child {
				border-right: none; /* Убираем правую границу у последнего элемента */
			}
			.navbar a:hover {
				background-color: #f5f5f5; /* Цвет фона при наведении аналогичен наведению на строки таблицы */
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
