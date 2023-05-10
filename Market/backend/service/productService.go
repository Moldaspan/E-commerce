package service

import (
	"github.com/Moldaspan/Market/Market/backend/database"
	"github.com/Moldaspan/Market/Market/backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.Db.Create(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func FilterItemsHandler(c *gin.Context) {
	// Read the query parameters
	nameFilter := c.Query("name")
	categoryFilter := c.Query("category")

	// Build the Gorm query
	query := database.Db.Model(&models.Product{})

	if nameFilter != "" {
		query = query.Where("name LIKE ?", "%"+nameFilter+"%")
	}

	if categoryFilter != "" {
		query = query.Where("category = ?", categoryFilter)
	}

	// Retrieve the filtered items from the database
	var items []models.Product
	if err := query.Find(&items).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Return the filtered items as a JSON response
	c.JSON(http.StatusOK, items)
}
