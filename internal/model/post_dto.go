package model

import "time"

// 定义请求的dto
type PostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostResponse struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
