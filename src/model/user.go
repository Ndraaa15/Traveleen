package model

import "gin/src/entity"

type UserRegister struct {
	Username string `json:"username" gorm:"unique" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	User  entity.User `json:"user"`
	Token string      `json:"token"`
}

type UserUpdate struct {
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Contact  string `json:"contact"`
	Region   string `json:"region"`
	Gender   bool   `json:"gender"`
	Birthday string `json:"birthday"`
}
