package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// TotalAccidentHandler обрабатывает запрос на получение общего количества ДТП и возвращает HTML-страницу.
func (s *Service) TotalAccidentHandler(w http.ResponseWriter, r *http.Request) {
	// Извлечение общего количества ДТП
	count, err := s.DB.GetAccidentCount()
	if err != nil {
		log.Printf("Ошибка при получении количества ДТП: %v", err)
		http.Error(w, "не удалось получить количество ДТП", http.StatusInternalServerError)
		return
	}

	// Формирование HTML-страницы с навигационной панелью
	tmpl := `
	<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Общее количество ДТП</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				margin: 20px;
			}
			h1 {
				color: #333;
			}
			.container {
				border: 1px solid #ddd;
				padding: 20px;
				border-radius: 5px;
				background-color: #f9f9f9;
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

		<div class="container">
			<h1>Общее количество ДТП</h1>
			<p>Всего ДТП: <strong>{{.TotalAccidents}}</strong></p>
		</div>
	</body>
	</html>
	`

	// Создание и выполнение шаблона
	t, err := template.New("totalAccidents").Parse(tmpl)
	if err != nil {
		log.Printf("Ошибка при разборе шаблона: %v", err)
		http.Error(w, "не удалось сгенерировать ответ", http.StatusInternalServerError)
		return
	}

	// Установка заголовков и возврат результата
	w.Header().Set("Content-Type", "text/html")
	if err := t.Execute(w, struct {
		TotalAccidents int
	}{TotalAccidents: count}); err != nil {
		log.Printf("Ошибка при выполнении шаблона: %v", err)
		http.Error(w, "не удалось сгенерировать ответ", http.StatusInternalServerError)
		return
	}
}
