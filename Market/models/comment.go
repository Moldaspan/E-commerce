package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Body      string `json:"body"`
	UserID    uint   `json:"user_id" binding:"required" gorm:"ForeignKey:User.ID" gorm:"constraint:OnDelete:CASCADE" `
	ProductID uint   `json:"product_id" binding:"required" gorm:"ForeignKey:Product.ID" gorm:"constraint:OnDelete:CASCADE" `
}
