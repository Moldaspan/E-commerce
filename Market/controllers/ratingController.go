package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RatingController struct {
	ratingService RatingServiceInterface
}

func NewRatingController() RatingController {
	return RatingController{ratingService: NewRatingService()}
}

func (rc RatingController) CreateRating(c *gin.Context) {
	var rating Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := rc.ratingService.CreateRating(&rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": rating})
}
