package entity

import (
	"gin/src/enum"
)

type User struct {
	ID           uint        `json:"id" gorm:"primarykey"`
	Username     string      `json:"username" gorm:"unique"`
	Email        string      `json:"email" gorm:"unique"`
	Password     string      `json:"-"`
	Wallet       string      `json:"wallet" `
	PhotoProfile string      `json:"photo_profile"`
	Contact      string      `json:"contact"`
	Region       string      `json:"region"`
	Gender       enum.Gender `json:"gender"`
	Birthday     string      `json:"birthday"`
}
