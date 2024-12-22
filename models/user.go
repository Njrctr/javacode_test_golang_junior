package models

type SignUpInput struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Id      int  `json:"-" db:"id"`
	IsAdmin bool `json:"-" db:"is_admin"`
	SignUpInput
}
