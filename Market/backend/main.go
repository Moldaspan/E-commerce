package main

import (
	"github.com/Moldaspan/E-commerce/backend/controllers"
	"github.com/Moldaspan/E-commerce/backend/database"
	"github.com/Moldaspan/E-commerce/backend/service"
	"github.com/gin-gonic/gin"
	"log"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "Market"
	ps     = "e!_sUltan747"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.LoadHTMLGlob("front/*.html")

	// Обработчики запросов без аутентификации
	router.GET("/register", service.ShowRegisterForm)
	router.POST("/register", service.RegisterHandler)
	router.GET("/login", service.ShowLoginForm)
	router.POST("/login", service.LoginHandler)
	router.POST("/products", controllers.ProductController{}.CreateProduct)
	router.GET("/products/:id", controllers.ProductController{}.GetProductByID)
	router.PUT("/products", controllers.ProductController{}.UpdateProduct)
	router.DELETE("/products/:id", controllers.ProductController{}.DeleteProduct)

	// Categories endpoints
	router.POST("/categories", controllers.CategoryController{}.CreateCategory)
	router.GET("/categories/:id", controllers.CategoryController{}.GetCategoryByID)
	router.PUT("/categories", controllers.CategoryController{}.UpdateCategory)
	router.DELETE("/categories/:id", controllers.CategoryController{}.DeleteCategory)
	//router.GET("/items", service.FilterItemsHandler)

	// Обработчики запросов с аутентификацией
	auth := router.Group("/")
	auth.Use(service.AuthMiddleware()) // использование middleware для проверки наличия JWT токена

	router.Run(":8888")
}
