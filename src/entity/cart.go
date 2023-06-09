package entity

import "gin/src/enum"

type Cart struct {
	ID          uint          `json:"id" gorm:"primarykey"`
	UserID      uint          `json:"user_id"`
	User        User          `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	CartProduct []CartProduct `json:"cart_product"`
	TotalPrice  float64       `json:"total_price"`
}

type CartProduct struct {
	CartID          uint             `json:"cart_id" gorm:"primarykey"`
	Cart            Cart             `json:"-" gorm:"foreignKey:CartID"`
	EcoImage        string           `json:"eco_image"`
	EcoID           uint             `json:"eco_id" gorm:"primarykey"`
	Ecotourism      Ecotourism       `json:"-" gorm:"foreignKey:EcoID"`
	EcoName         string           `json:"eco_name"`
	EcoLocation     string           `json:"eco_location"`
	EcoCategory     enum.EcoCategory `json:"eco_category"`
	Quantity        uint             `json:"quantity"`
	Price           float64          `json:"price"`
	PricePerProduct float64          `json:"price_per_product"`
}
