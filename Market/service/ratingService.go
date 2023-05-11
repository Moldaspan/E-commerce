package service

import (
	"github.com/Moldaspan/E-commerce/models"
	"github.com/Moldaspan/E-commerce/repositories"
)

type RatingServiceInterface interface {
	CreateRating(rating *models.Rating) error
}

type RatingServiceV1 struct {
	ratingRepos RatingServiceInterface
}

func NewRatingService() RatingServiceV1 {
	return RatingServiceV1{ratingRepos: repositories.NewRatingRepository()}
}

func (r RatingServiceV1) CreateRating(rating *models.Rating) error {
	return r.ratingRepos.CreateRating(rating)
}
