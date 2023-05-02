package entity

import "gorm.io/datatypes"

type Comment struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	Date         string         `json:"date"`
	UserID       uint           `json:"user_id"`
	User         User           `json:"user" foreignKey:"UserID"`
	Rating       float64        `json:"rating"`
	Body         string         `json:"body"`
	Thumbnail    datatypes.JSON `json:"thumbnail" gorm:"type:text"`
	EcotourismID uint           `json:"ecotourism_id"`
	Ecotourism   Ecotourism     `json:"-" foreignKey:"EcotourismID"`
}
