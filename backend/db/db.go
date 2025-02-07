package db

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v5/pgxpool"
    "your_project/models"
    "github.com/jackc/pgconn"
)

type Database struct {
    pool *pgxpool.Pool
}

func NewDatabase(connString string) (*Database, error) {
    pool, err := pgxpool.New(context.Background(), connString)
    if err != nil {
        return nil, fmt.Errorf("unable to create connection pool: %w", err)
    }

    db := &Database{pool: pool}
    if err := db.Init(); err != nil {
        return nil, err
    }
    return db, nil
}

func (d *Database) Init() error {
    query := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) UNIQUE NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    )`
    
    _, err := d.pool.Exec(context.Background(), query)
    return err
}

func (d *Database) CreateUser(user *models.User) error {
    query := `INSERT INTO users (username, email, password_hash) 
              VALUES ($1, $2, $3) RETURNING id, created_at`
    
    err := d.pool.QueryRow(context.Background(), query, 
        user.Username, 
        user.Email, 
        user.PasswordHash,
    ).Scan(&user.ID, &user.CreatedAt)

    if pgErr, ok := err.(*pgconn.PgError); ok {
        if pgErr.Code == "23505" {
            return fmt.Errorf("user with this email or username already exists")
        }
    }
    
    return err
}

func (d *Database) GetUsers() ([]models.User, error) {
    query := `SELECT id, username, email, created_at FROM users`
    rows, err := d.pool.Query(context.Background(), query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var u models.User
        err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt)
        if err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    
    return users, nil
}
