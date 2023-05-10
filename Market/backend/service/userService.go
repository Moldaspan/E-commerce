package service

import (
	"github.com/Moldaspan/E-commerce/backend/database"
	"github.com/Moldaspan/E-commerce/backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ShowRegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"ErrorMessage": "",
	})
}

func RegisterHandler(c *gin.Context) {
	// Получение данных из формы
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Проверка наличия соединения с базой данных
	if database.Db == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Проверка, существует ли пользователь с таким же email
	var existingUser models.User
	result := database.Db.Where("email = ?", email).First(&existingUser)
	if result.Error == nil {
		// Пользователь уже существует
		c.HTML(http.StatusOK, "register.html", gin.H{
			"ErrorMessage": "Пользователь с таким email уже зарегистрирован",
		})
		return
	}

	// Создание нового пользователя
	user := models.User{Firstname: firstname, LastName: lastname, Email: email, Password: password}
	result = database.Db.Create(&user)
	if result.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.String(http.StatusOK, "User %s created successfully!", email)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		// Проверка валидности JWT токена
		token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return database.JwtSecret, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Добавление декодированного токена в контекст Gin
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			// Проверяем, является ли пользователь админом
			var admin models.Admin
			result := database.Db.Where("email = ?", email).First(&admin)
			if result.Error == nil {
				// Если пользователь - администратор, то создаем новый токен с другим секретным ключом
				newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"email": email,
					"admin": true,
				})
				tokenString, err := newToken.SignedString([]byte("admin-secret-key"))
				if err != nil {
					c.AbortWithStatus(http.StatusInternalServerError)
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"token": tokenString,
				})
				return
			}

			// Если пользователь - обычный пользователь, то возвращаем успешный ответ
			c.JSON(http.StatusOK, gin.H{
				"message": "You are authorized!",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
		}
	}
}

// Хэндлер для страницы авторизации
func ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"ErrorMessage": "",
	})
}

func LoginHandler(c *gin.Context) {
	// Получение данных из формы
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Проверка наличия соединения с базой данных
	if database.Db == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Проверка, существует ли пользователь с таким email и паролем
	var user models.User
	result := database.Db.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		// Пользователь не найден
		c.HTML(http.StatusOK, "login.html", gin.H{
			"ErrorMessage": "Неверный email или пароль",
		})
		return
	}

	// Создание токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	println(tokenString)
	// Установка токена в куки
	c.SetCookie("token", tokenString, int(time.Hour.Seconds()), "/", "", false, true)
	// Перенаправление на главную страницу
	//c.Redirect(http.StatusFound, "/register")
}
