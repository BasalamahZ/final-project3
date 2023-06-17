package dto

type UserRequest struct {
	Fullname string  `json:"full_name" binding:"required"`
	Email    string  `json:"email" gorm:"unique" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=6"`
	Role     string `json:"role" gorm:"default:'Member'"`
}

type LoginRequest struct {
	Email    string `json:"email" gorm:"unique" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type EditUser struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
}
