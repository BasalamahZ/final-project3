package model

import "time"

type User struct {
	Id        int     `json:"id" gorm:"primaryKey;UNIQUE"`
	Fullname  string    `json:"full_name" binding:"required"`
	Email     string    `json:"email" gorm:"unique" binding:"required,email"`
	Password  string    `json:"password" binding:"required,min=6"`
	Role      string    `json:"role" binding:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}
