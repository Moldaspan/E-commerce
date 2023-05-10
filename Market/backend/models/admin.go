package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Email    string `gorm:"not null unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
