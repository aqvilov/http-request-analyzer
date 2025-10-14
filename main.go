package main

import (
	"fmt"
	"http-request-analyzer/analyzer"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All good")
}

func main() {
	// 1. Сначала объявляем ВСЕ маршруты
	http.HandleFunc("/reg", test)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("✅ Запрос получен!")
		fmt.Fprintf(w, "Данные о сервере выведены в консоль!")
		analyzer.PrintRequest(r)
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			name := r.FormValue("name")
			email := r.FormValue("email")
			fmt.Fprintf(w, "✅ Данные получены!\nИмя: %s\nEmail: %s", name, email)
		} else {
			fmt.Fprintf(w, `
				<h1>Тест POST запроса</h1>
				<form method="POST" action="/submit">
					<input type="text" name="name" placeholder="Имя"><br>
					<input type="email" name="email" placeholder="Email"><br>
					<button type="submit">Отправить</button>
				</form>
			`)
		}
		analyzer.PrintRequest(r)
	})

	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "POST" {
			fmt.Fprintf(w, `{"status": "user created", "method": "POST"}`)
		} else if r.Method == "GET" {
			fmt.Fprintf(w, `{"users": ["john", "alice"], "method": "GET"}`)
		} else {
			fmt.Fprintf(w, `{"error": "method not allowed"}`)
		}
		analyzer.PrintRequest(r)
	})

	// 2. Потом выводим информацию
	fmt.Println("🚀 Сервер запущен на http://localhost:8081")
	fmt.Println("📝 Тестовые endpoints:")
	fmt.Println("   GET  http://localhost:8081/")
	fmt.Println("   GET  http://localhost:8081/submit")
	fmt.Println("   POST http://localhost:8081/submit")
	fmt.Println("   POST http://localhost:8081/api/users")

	// 3. И только в конце запускаем сервер
	http.ListenAndServe(":8081", nil)
}
