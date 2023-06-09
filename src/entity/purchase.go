package entity

import "gin/src/enum"

type Purchase struct {
	ID          uint             `json:"id" gorm:"primarykey, unique"`
	Date        string           `json:"date"`
	Time        string           `json:"time"`
	Place       string           `json:"place"`
	Quantity    uint             `json:"quantity"`
	TotalPrice  float64          `json:"total_price"`
	Code        string           `json:"code"`
	Status      enum.Status      `json:"status"`
	PayCategory enum.PayCategory `json:"pay_category"`
	EcoName     string           `json:"eco_name"`
	EcoCategory enum.EcoCategory `json:"eco_category"`
	EcoLocation string           `json:"eco_location"`
	EcoImage    string           `json:"eco_image"`
	UserID      uint             `json:"user_id"`
	User        User             `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
