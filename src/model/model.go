package model

import (
	"gin/src/entity"

	"gorm.io/datatypes"
)

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

type UserComment struct {
	Rating    float64        `json:"rating" binding:"required"`
	Body      string         `json:"body"  binding:"required"`
	Thumbnail datatypes.JSON `json:"thumbnail" gorm:"type:json"`
}

type CartProduct struct {
	Quantity uint `json:"quantity" binding:"required"`
}

type PaymentType struct {
	IsOnlinePayment bool `json:"is_online_payment"`
}

type PurchaseResponse struct {
	IsSuccess   string `json:"is_success"`
	PaymentType string `json:"payment_type"`
	URL         string `json:"url"`
}

type ExchangeTrash struct {
	Category string  `json:"category"`
	Mass     float64 `json:"mass"`
}
