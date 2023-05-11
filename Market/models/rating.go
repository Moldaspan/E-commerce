package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	UserID    uint `json:"user_id" binding:"required" gorm:"ForeignKey:User.ID; constraint:OnDelete:CASCADE;unique" `
	ProductID uint `json:"product_id" binding:"required" gorm:"ForeignKey:Product.ID; constraint:OnDelete:CASCADE" `
	Value     uint `json:"value"  binding:"required"`
}
