package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/your_project/internal/db"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, sqlDB *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Используем имя поля "username" (см. index.html)
	username := r.FormValue("username")
	email := r.FormValue("email")

	if username == "" || email == "" {
		http.Error(w, "Fill in all the fields", http.StatusBadRequest)
		return
	}

	// Регистрируем пользователя без пароля (password_hash оставляем пустым)
	query := "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id, created_at"
	var id int
	var createdAt string
	if err := sqlDB.QueryRow(query, username, email, "").Scan(&id, &createdAt); err != nil {
		log.Println("Error when inserting into the database:", err)
		http.Error(w, "Registration error", http.StatusInternalServerError)
		return
	}

	// Получаем список пользователей для отображения на доске
	users, err := db.GetUsers(sqlDB)
	if err != nil {
		http.Error(w, "Error retrieving users", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles(filepath.Join("frontend", "templates", "board.html"))
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, users)
}
