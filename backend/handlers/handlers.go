package handlers

import (
    "net/http"
    "encoding/json"
    "your_project/models"
    "your_project/db"
)

type App struct {
    db *db.Database
}

func NewApp(database *db.Database) *App {
    return &App{db: database}
}

func (a *App) RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req struct {
        Username string `json:"username"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if req.Username == "" || req.Email == "" || req.Password == "" {
        http.Error(w, "All fields are required", http.StatusBadRequest)
        return
    }

    user := models.User{
        Username: req.Username,
        Email:    req.Email,
    }

    if err := user.SetPassword(req.Password); err != nil {
        http.Error(w, "Error processing password", http.StatusInternalServerError)
        return
    }

    if err := a.db.CreateUser(&user); err != nil {
        http.Error(w, err.Error(), http.StatusConflict)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func (a *App) UsersHandler(w http.ResponseWriter, r *http.Request) {
    users, err := a.db.GetUsers()
    if err != nil {
        http.Error(w, "Error retrieving users", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
