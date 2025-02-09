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
		http.Error(w, "Only POST-reauests", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if name == "" || email == "" {
		http.Error(w, "Fill in all the fields", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	if err != nil {
		log.Println("Error when inserting into the database:", err)
		http.Error(w, "Registartion error", http.StatusInternalServerError)
		return
	}


	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		log.Println("Error receiving users:", err)
		http.Error(w, "Error of loading list of users", http.StatusInternalServerError)
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


	tmpl, err := template.ParseFiles("frontend/templates/board.html")
	if err != nil {
		http.Error(w, "Error downloading page", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, users)
}
