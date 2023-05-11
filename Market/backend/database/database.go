package database

import (
	"fmt"
	"github.com/Moldaspan/E-commerce/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var (
	Db             *gorm.DB
	JwtSecret      = []byte("my-secret-key")
	TokenExpiresIn = time.Hour * 24 // JWT token expiration time
)

func InitDB() (*gorm.DB, error) {
	// Set up database connection
	dsn := "host=db port = 5432 user = postgres dbname = bookstore password = Ayef1407_ sslmode = disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	err = migrateDB(db)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return db, nil
}
func migrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Admin{}, models.User{}, models.Category{}, models.Product{})
	if err != nil {
		return fmt.Errorf("failed to migrate books table: %v", err)
	}

	return nil
}
