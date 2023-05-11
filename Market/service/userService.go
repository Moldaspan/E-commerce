package service

import (
	"github.com/Moldaspan/E-commerce/models"
	"github.com/Moldaspan/E-commerce/repositories"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	Id        uint
	UserType  string
	jwt.StandardClaims
}

type InvalidCredentialsError struct {
	msg string
}

func (e *InvalidCredentialsError) Error() string {
	return e.msg
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

type UserServiceInterface interface {
	CreateUser(user *models.User) error
	LogIn(user *models.User) (string, string, error)
	GetUserByID(id uint) *models.User
}

type UserServiceV1 struct {
	userRepos repositories.UserReposInterface
}

func (u UserServiceV1) CreateUser(user *models.User) error {
	user.Password = HashPassword(user.Password)
	return u.userRepos.CreateUser(user)
}

func (u UserServiceV1) GetUserByID(id uint) *models.User {
	return u.userRepos.GetUserByID(id)
}

func (u UserServiceV1) LogIn(user *models.User) (string, string, error) {
	givenpwd := user.Password
	user, err := u.userRepos.GetUser(user.Email)
	if err != nil {
		return "", "", err
	}

	valid, msg := ComparePasswords(user.Password, givenpwd)

	if valid == false {
		return "", "", &InvalidCredentialsError{msg: msg}
	}

	access, refresh, err := TokenGenerator(user)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil

}

func NewUserService() UserServiceV1 {
	return UserServiceV1{userRepos: repositories.NewUserRepos()}
}

func TokenGenerator(user *models.User) (string, string, error) {
	claims := &SignedDetails{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Id:        user.ID,
		UserType:  user.UserType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panicln(err)
		return "", "", err
	}
	return token, refreshToken, err
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func ComparePasswords(hashedpwd string, givenpwd string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpwd), []byte(givenpwd))
	valid := true
	msg := ""
	if err != nil {
		msg = "Email or password is incorrect"
		valid = false
	}
	return valid, msg
}
