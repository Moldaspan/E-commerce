package database

import (
	"gorm.io/gorm"
	"time"
)

var (
	Db             *gorm.DB
	JwtSecret      = []byte("my-secret-key")
	TokenExpiresIn = time.Hour * 24 // JWT token expiration time
)
