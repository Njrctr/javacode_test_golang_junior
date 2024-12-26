package models

type SignUpInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	SignUpInput
	Id      int  `json:"-" db:"id"`
	IsAdmin bool `json:"-" db:"is_admin"`
}
