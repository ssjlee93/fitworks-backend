package entities

type Album struct {
	ID     int     `gorm:"primaryKey"`
	Title  string  `gorm:"type:varchar(255)"`
	Artist string  `gorm:"type:varchar(255)"`
	Price  float64 `gorm:"type:float"`
}
