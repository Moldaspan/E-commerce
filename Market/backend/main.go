package main


import (
	"fmt"
	"github.com/Moldaspan/E-commerce/backend/database"
	"github.com/Moldaspan/E-commerce/backend/models"
	"github.com/Moldaspan/E-commerce/backend/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//
func(){

}
const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "Market"
	ps     = "e!_sUltan747"
)

func main() {
	// Создание подключения к базе данных
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, ps, dbname)
	var err error
	database.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Проверка соединения с базой данных
	if err := database.Db.AutoMigrate(&models.User{}, &models.Admin{}); err != nil {
		panic("failed to migrate database schema")
	}

	router := gin.Default()
	router.LoadHTMLGlob("front/*.html")

	// Обработчики запросов без аутентификации
	router.GET("/register", service.ShowRegisterForm)
	router.POST("/register", service.RegisterHandler)
	router.GET("/login", service.ShowLoginForm)
	router.POST("/login", service.LoginHandler)
	router.GET("/items", service.FilterItemsHandler)

	// Обработчики запросов с аутентификацией
	auth := router.Group("/")
	auth.Use(service.AuthMiddleware()) // использование middleware для проверки наличия JWT токена

	router.Run(":8888")
}
