package model

type User struct {
	ID       int `gorm:"primary_key;auto_increment"`
	Name     string `gorm:"size:100;not null" json:"name"`
	Email    string `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
}