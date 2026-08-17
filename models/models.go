package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Password  string    `json:"-" gorm:"uniqueIndex;not null"`
	EmailID   string    `json:"email" gorm:"uniqueIndex;not null"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	EmailID  string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"gte=0,lte=150"`
}

type LoginUser struct {
	EmailID  string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Tweet struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	User      User      `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTweet struct {
	Content string `json:"content" binding:"required,max=280"`
}
