package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null;size:30" json:"first_name"`
	LastName  string `gorm:"not null;size:30" json:"last_name"`
	Password  string `gorm:"not null;size:255" json:"password"`
	Email     string `gorm:"not null;size:255;unique" json:"email"`
	UserType  string `gorm:"not null;size:20" json:"user_type"`
}
