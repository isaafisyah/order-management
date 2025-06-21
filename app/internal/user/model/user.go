package model

type User struct {
	ID       int `gorm:"primary_key;auto_increment"`
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;not null;uniqueIndex"`
	Password string `gorm:"size:255;not null;"`
	CreatedAt string `gorm:"autoCreateTime"`
}