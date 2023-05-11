package service

type RatingServiceInterface interface {
	CreateRating(rating *Rating) error
}

type RatingServiceV1 struct {
	ratingRepos RatingServiceInterface
}

func NewRatingService() RatingServiceV1 {
	return RatingServiceV1{ratingRepos: NewRatingRepository()}
}

func (r RatingServiceV1) CreateRating(rating *Rating) error {
	return r.ratingRepos.CreateRating(rating)
}
