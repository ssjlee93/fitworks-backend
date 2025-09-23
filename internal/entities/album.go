package entities

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Title  string  `gorm:"type:varchar(255)"`
	Artist string  `gorm:"type:varchar(255)"`
	Price  float64 `gorm:"type:float"`
}
