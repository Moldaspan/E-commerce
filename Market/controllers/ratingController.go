package controllers

import (
	"github.com/Moldaspan/E-commerce/models"
	"github.com/Moldaspan/E-commerce/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RatingController struct {
	ratingService service.RatingServiceInterface
}

func NewRatingController() RatingController {
	return RatingController{ratingService: service.NewRatingService()}
}

func (rc RatingController) CreateRating(c *gin.Context) {
	var rating models.Rating
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
