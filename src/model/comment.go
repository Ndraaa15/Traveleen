package model

type Comment struct {
	Rating    float64  `json:"rating"`
	Body      string   `json:"body"`
	Thumbnail []string `json:"thumbnail" gorm:"type:text"`
}
