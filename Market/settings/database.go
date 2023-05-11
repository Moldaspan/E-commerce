package settings

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "ecommercegolang"
	ps     = "Ayef1407_"
)

func DbSetup() (*gorm.DB, error) {
	godotenv.Load(".env")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, ps, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
