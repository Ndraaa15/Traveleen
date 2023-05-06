package entity

type Article struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Thumbnail string `json:"thumbnail"`
	Date      string `json:"date"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"-" gorm:"foreignKey:user_id"`
}
