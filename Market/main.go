package main

import (
	"github.com/Moldaspan/E-commerce/controllers"
	"github.com/Moldaspan/E-commerce/models"
	"github.com/Moldaspan/E-commerce/settings"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	userController := controllers.NewUserController()
	pc := controllers.NewProductController()
	cc := controllers.NewCategoryController()
	oc := controllers.NewOrderController()

	db, _ := settings.DbSetup()
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.Order{})

	router.POST("/users", userController.CreateUser)
	router.GET("/users/:user_id", userController.GetUserById)
	router.POST("/users/token", userController.LogIn)

	router.POST("/orders", oc.CreateOrder)
	router.GET("/orders/:id", oc.GetOrderById)
	router.PUT("/orders/:id", oc.UpdateOrder)
	router.GET("/users/:user_id/orders", oc.GetOrders)
	router.DELETE("/orders/:id", oc.DeleteOrder)

	router.POST("/products", pc.CreateProduct)
	router.GET("/products/", pc.SearchByName)
	router.GET("/products/filter", pc.SearchByPriceRange)
	router.GET("/products/:id", pc.GetProductByID)
	router.PUT("/products", pc.UpdateProduct)
	router.DELETE("/products/:id", pc.DeleteProduct)

	// Categories endpoints
	router.POST("/categories", cc.CreateCategory)
	router.GET("/categories/:id", cc.GetCategoryByID)
	router.PUT("/categories", cc.UpdateCategory)
	router.DELETE("/categories/:id", cc.DeleteCategory)

	log.Fatal(http.ListenAndServe(":8080", router))
}
