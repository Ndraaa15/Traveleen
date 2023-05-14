package entity

import (
	"gin/src/enum"
)

type User struct {
	ID           uint        `json:"id" gorm:"primarykey"`
	Username     string      `json:"username" gorm:"unique"`
	Email        string      `json:"email" gorm:"unique" binding:"email"`
	Password     string      `json:"-"`
	Wallet       float64     `json:"wallet" `
	PhotoProfile string      `json:"photo_profile" binding:"e164"`
	Contact      string      `json:"contact" gorm:"unique"`
	Region       string      `json:"region"`
	Gender       enum.Gender `json:"gender"`
	Birthday     string      `json:"birthday"`
}
