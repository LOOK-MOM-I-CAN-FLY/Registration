package main

import (
	"log"
	"net/http"

	"github.com/your_project/internal/db"
	"github.com/your_project/internal/handlers"
)

func main() {
	// Инициализация БД
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Ошибка инициализации БД:", err)
	}
	defer database.Close()

	// Настройка маршрутов
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterUser(w, r, database)
	})

	// Запуск сервера
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
