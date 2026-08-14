package models

import (
	"time"
)

type User struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Password  string    `json:"-" gorm:"uniqueIndex;not null"`
	EmailID  string    `json:"email" gorm:"uniqueIndex;not null"`
	Age     int       `json:"age"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateUser struct {
	Name    string `json:"name"`
	Password string `json:"password" `
	EmailID string `json:"email"`
	Age    int    `json:"age"`
}

type RegisterUser struct {
	Name    string `json:"name"`
	Password string `json:"password" `
	EmailID string `json:"email"`
	Age   int    `json:"age"`
}

type LoginUser struct {
	EmailID string `json:"email"`
	Password string `json:"password" `
}
