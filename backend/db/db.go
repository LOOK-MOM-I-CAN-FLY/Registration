package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
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

	// Создание таблицы, если её нет
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL
	)`)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}

	log.Println("Подключение к БД успешно")
	return db, nil
}
