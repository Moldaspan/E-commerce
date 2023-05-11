package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title      string `gorm:"size: 150;unique"`
	Body       string `json:"body"`
	MainImage  string `json:"main_image"`
	IsActive   bool   `json:"is_active"`
	IsTop      bool   `json:"is_top"`
	UserID     uint   `json:"user_id" binding:"required" gorm:"ForeignKey:User.ID" gorm:"constraint:OnDelete:CASCADE" `
	CategoryID uint   `json:"category_id" binding:"required" gorm:"ForeignKey:Category.ID" gorm:"constraint:OnDelete:CASCADE" `

	Price float64
}
