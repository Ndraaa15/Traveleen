package model

import (
	"gin/src/enum"

	"gorm.io/datatypes"
)

type PostEcotourisms struct {
	Region          string           `json:"region"`
	Category        enum.EcoCategory `json:"category"`
	Thumbnail       datatypes.JSON   `json:"thumbnail" gorm:"type:json"`
	Name            string           `json:"name"`
	Rating          float64          `json:"rating"`
	TotalRatings    uint             `json:"total_ratings"`
	Price           float64          `json:"price"`
	Description     string           `json:"description"`
	OperationalTime datatypes.JSON   `json:"operational_time" gorm:"type:json"`
	Maps            string           `json:"maps"`
}
