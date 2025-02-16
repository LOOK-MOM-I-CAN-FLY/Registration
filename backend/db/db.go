package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/your_project/internal/models"
)

const dsn = "host=localhost port=5432 user=postgres password=yourpassword dbname=yourdb sslmode=disable"

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Создаём таблицу, если её нет
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password_hash TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal("Error creating the table:", err)
	}

	log.Println("Connection to the database is successful")
	return db, nil
}

func CreateUser(sqlDB *sql.DB, user *models.User) error {
	query := "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id, created_at"
	return sqlDB.QueryRow(query, user.Username, user.Email, user.PasswordHash).Scan(&user.ID, &user.CreatedAt)
}

func GetUsers(sqlDB *sql.DB) ([]models.User, error) {
	rows, err := sqlDB.Query("SELECT id, username, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt); err != nil {
			continue
		}
		users = append(users, u)
	}
	return users, nil
}
