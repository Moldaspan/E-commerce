package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Firstname string `gorm:"not null" json:"firstname"`
	LastName  string `gorm:"not null" json:"lastName"`
	Email     string `gorm:"not null unique" json:"email"`
	Password  string `gorm:"not null" json:"password"`
}
