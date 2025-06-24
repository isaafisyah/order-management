package model

type Product struct {
	ID          int    `gorm:"primary_key;auto_increment"`
	Name        string `gorm:"size:100;not null"`
	CategoryId  int    `gorm:"not null"`
	Price       int    `gorm:"not null"`
	Stock 		int    `gorm:"not null"`
	CreatedAt   string `gorm:"autoCreateTime"`
	UpdatedAt   string `gorm:"autoUpdateTime"`
}