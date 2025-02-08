package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/your_project/internal/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST-запросы", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if name == "" || email == "" {
		http.Error(w, "Заполните все поля", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	if err != nil {
		log.Println("Ошибка при вставке в БД:", err)
		http.Error(w, "Ошибка регистрации", http.StatusInternalServerError)
		return
	}

	// Получение всех пользователей для отображения
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Println("Ошибка получения пользователей:", err)
		http.Error(w, "Ошибка загрузки списка пользователей", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Println("Ошибка сканирования:", err)
			continue
		}
		users = append(users, user)
	}

	// Отображение страницы с пользователями
	tmpl, err := template.ParseFiles("frontend/templates/board.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки страницы", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, users)
}
