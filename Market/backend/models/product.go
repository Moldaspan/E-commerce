package models

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Price       uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	amount      int    `gorm:"not null"`
}
