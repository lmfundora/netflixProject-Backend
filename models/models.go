package models

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Name string `json:"name" gorm:"text;not null;unique;default:null"`
	Phone string `json:"phone" gorm:"text;not null;unique;default:null"`
	Email string  `json:"email" gorm:"text;not null;unique;default:null"`
	Password string `json:"password" gorm:"text;not null;default:null"`
}

type LoginRequest struct {
	Email	string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}


