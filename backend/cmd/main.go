package main

import (
	"html/template"
	"net/http"
	"sync"
)

type User struct {
	FirstName string
	LastName  string
}

var (
	users     []User
	usersLock sync.Mutex
	templates *template.Template
)

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", registrationHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/users", usersHandler)

	// Запуск сервера на порте 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		firstName := r.FormValue("first_name")
		lastName := r.FormValue("last_name")

		// Добавляем нового пользователя в список с блокировкой
		usersLock.Lock()
		users = append(users, User{FirstName: firstName, LastName: lastName})
		usersLock.Unlock()

		// Перенаправляем на страницу со списком пользователей
		http.Redirect(w, r, "/users", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	usersLock.Lock()
	defer usersLock.Unlock()
	templates.ExecuteTemplate(w, "users.html", users)
}
