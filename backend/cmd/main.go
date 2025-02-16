package main

import (
	"database/sql"
	"log"
	"net/http"
	"path/filepath"

	_ "github.com/lib/pq"
	"github.com/your_project/internal/db"
	"github.com/your_project/internal/handlers"
)

func main() {
	// Инициализация БД
	sqlDB, err := db.InitDB()
	if err != nil {
		log.Fatal("Ошибка инициализации БД:", err)
	}
	defer sqlDB.Close()

	// Раздача статических файлов (CSS, JS)
	fs := http.FileServer(http.Dir(filepath.Join("frontend", "static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Маршруты для HTML-страниц
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterUser(w, r, sqlDB)
	})

	// API-эндпоинты
	http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterAPI(w, r, sqlDB)
	})
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.UsersAPI(w, r, sqlDB)
	})

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
