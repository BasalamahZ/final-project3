package dto

import (
	"time"
)

type CategoryRequest struct {
	Type string `json:"type" binding:"required"`
}

type CategoryResponse struct {
	Id        int64      `json:"id"`
	Type      string     `json:"type"`
	Tasks     []TaskData `json:"Tasks"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type TaskData struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
