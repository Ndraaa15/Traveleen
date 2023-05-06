package entity

import "gin/src/enum"

type Booking struct {
	ID         uint        `json:"id" gorm:"primarykey"`
	Date       string      `json:"date"`
	Place      string      `json:"place"` //Pantai Kuta, Bali
	Quantity   int         `json:"quantity"`
	TotalPrice string      `json:"total_price"`
	Status     enum.Status `json:"status"`
	Code       string      `json:"code"`
}
