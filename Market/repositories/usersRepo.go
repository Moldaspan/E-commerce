package repositories

import (
	"github.com/Moldaspan/E-commerce/models"
	"github.com/Moldaspan/E-commerce/settings"
	"gorm.io/gorm"
	"log"
)

type UserReposInterface interface {
	CreateUser(u *models.User) error
	GetUser(email string) (*models.User, error)
	GetUserByID(id uint) *models.User
}

type UserReposV1 struct {
	DB *gorm.DB
}

func NewUserRepos() UserReposInterface {
	db, err := settings.DbSetup()
	if err != nil {
		log.Fatalf("Error %s", err)
		return &UserReposV1{}
	}
	return &UserReposV1{DB: db}
}

func (u UserReposV1) CreateUser(user *models.User) error {
	return u.DB.Create(&user).Error
}

func (u UserReposV1) GetUser(email string) (*models.User, error) {
	var user models.User
	res := u.DB.Where("email = ?", email).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (u UserReposV1) GetUserByID(id uint) *models.User {
	var user models.User
	u.DB.Where("id = ?", id).First(&user)
	return &user
}
