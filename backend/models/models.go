package models

import (
    "time"
    "errors"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    ID          int       `json:"id"`
    Username    string    `json:"username"`
    Email       string    `json:"email"`
    PasswordHash string   `json:"-"`
    CreatedAt   time.Time `json:"created_at"`
}

func (u *User) SetPassword(password string) error {
    if len(password) == 0 {
        return errors.New("password cannot be empty")
    }
    bytePassword := []byte(password)
    hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.PasswordHash = string(hashedPassword)
    return nil
}

func (u *User) CheckPassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}
