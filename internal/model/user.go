package model

import "time"

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     *string   `json:"email"` // 可为空
	CreatedAt time.Time `json:"created_at"`
}

type UserRegisterDto struct {
	Name     string  `json:"name" binding:"required"`
	Password string  `json:"password" binding:"required""`
	Email    *string `json:"email"` // 可选
}

type UserLoginDto struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
