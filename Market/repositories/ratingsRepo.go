package repositories

import (
	"github.com/Moldaspan/E-commerce/models"
	"github.com/Moldaspan/E-commerce/settings"
	"gorm.io/gorm"
	"log"
)

type RatingRepositoryInterface interface {
	CreateRating(*models.Rating) error
}

type RatingRepositoryV1 struct {
	DB *gorm.DB
}

func NewRatingRepository() *RatingRepositoryV1 {
	db, err := settings.DbSetup()
	if err != nil {
		log.Fatal(err)
		return &RatingRepositoryV1{}
	}
	return &RatingRepositoryV1{DB: db}
}

func (r *RatingRepositoryV1) CreateRating(rating *models.Rating) error {
	return r.DB.Create(rating).Error
}
