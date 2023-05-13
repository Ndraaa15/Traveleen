package entity

import (
	"gin/src/enum"
)

type Trash struct {
	ID            uint               `gorm:"primarykey"`
	Date          string             `json:"date"`
	Category      enum.TrashCategory `json:"category"`
	Location      string             `json:"location"`
	Mass          float64            `json:"mass"`
	Code          string             `json:"code"`
	Status        enum.Status        `json:"status"`
	ExchangeTotal float64            `json:"exchange_totals"`
	UserID        uint               `json:"user_id"`
	User          User               `gorm:"foreignKey:UserID"`
}
