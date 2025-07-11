package model


type Category struct {
	ID   int    `gorm:"primary_key;auto_increment"`
	Name string `gorm:"size:100;not null;uniqueIndex"`
}