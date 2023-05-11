package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID uint `json:"product_id" binding:"required" gorm:"ForeignKey:Product.ID"`
	UserID    uint `json:"user_id" binding:"required" gorm:"ForeignKey:User.ID"`
	Status    string
}
