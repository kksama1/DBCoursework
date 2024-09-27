package handlers

import (
	"html/template"
	"net/http"
)

//func (s *Service) GetAccidentReportByDayAndTimeHandler(w http.ResponseWriter, r *http.Request) {
//	// Извлечение отчёта по дням недели и времени суток
//	report, err := s.DB.GetAccidentReportByDayAndTime()
//	if err != nil {
//		log.Printf("Error fetching accident report: %v", err)
//		http.Error(w, "failed to fetch accident report", http.StatusInternalServerError)
//		return
//	}
//
//	// Проверка, найден ли отчёт
//	if len(report) == 0 {
//		http.Error(w, "no data found", http.StatusNotFound)
//		return
//	}
//
//	// Возврат результата в виде JSON
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(report)
//}

func (s *Service) GetAccidentReportByDayAndTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ДТП по дням недели и времени суток из базы данных
	accidents, err := s.DB.GetAccidentReportByDayAndTime()
	if err != nil {
		http.Error(w, "Failed to fetch accidents", http.StatusInternalServerError)
		return
	}

	// Определяем HTML-шаблон
	tmpl := `
	<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>ДТП по дням и времени</title>
		<style>
			table {
				width: 100%;
				border-collapse: collapse;
			}
			th, td {
				border: 1px solid #ddd;
				padding: 8px;
			}
			th {
				background-color: #f2f2f2;
			}
			tr:hover {
				background-color: #f5f5f5;
			}
		</style>
	</head>
	<body>
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
