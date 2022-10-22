package services

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-server/models"
	"net/http"
	"os"
	"time"
)

// RegisterUser Регистрация
func RegisterUser(c *gin.Context) {
	var input models.UserRegisterData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	var userCheck models.User

	if err := models.DB.Where("email=?", input.Email).First(&userCheck).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"msg": "Электронная почта уже занята"})
		return
	}

	user := models.User{Email: input.Email, Password: MD5(input.Password), FirstName: input.FirstName, LastName: input.LastName, Role: input.Role, EmailCheck: false}
	models.DB.Create(&user)

	emailCheck := models.EmailCheck{ID: user.ID, UUID: uuid.New().String()}
	models.DB.Create(&emailCheck)

	SendEmailUUID(user.Email, emailCheck.UUID)

	c.JSON(http.StatusOK, gin.H{"msg": "Мы отправили письмо с подтверждением на вашу электронную почту. Вы сможете зайти на аккант только после подтверждения."})
}

// LoginUser Авторизация
func LoginUser(c *gin.Context) {
	var input models.UserLoginData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	var user models.User

	if err := models.DB.Where("email=?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Пользователь не найден"})
		return
	}

	if user.Password != MD5(input.Password) {
		c.JSON(http.StatusForbidden, gin.H{"msg": "Неверный пароль"})
		return
	}

	if !user.EmailCheck {
		c.JSON(http.StatusForbidden, gin.H{"msg": "Подтвердите адрес электронной почты"})
		return
	}

	token := CreateToken(user)
	//tokens := models.Token{ID: user.ID, Access: CreateToken(user), Refresh: CreateTokenRefresh()}

	//if err := models.DB.Where("id=?", user.ID).First(&token).Error; err != nil {
	//	models.DB.Create(&tokens)
	//	c.SetCookie("refresh_token", tokens.Refresh, 60*60*24*30, "/", os.Getenv("cookie_http"), false, true) // if https: secure = true
	//	c.JSON(http.StatusOK, gin.H{"access": tokens.Access})
	//	return
	//}
	//models.DB.Model(&token).Updates(tokens)
	//c.SetCookie("refresh_token", tokens.Refresh, 60*60*24*30, "/", os.Getenv("cookie_http"), false, true) // if https: secure = true
	c.JSON(http.StatusOK, gin.H{"access": token})
}

// CreateToken Создание JWT
func CreateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": user.ID,
		"email":  user.Email,
		"exp":    time.Now().Add(time.Hour * 24 * 2).Unix(),
	})

	jwtToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return jwtToken
}

// CreateTokenRefresh Создание Refresh токена
//func CreateTokenRefresh() string {
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//		"exp": time.Now().Add(time.Hour * 7200).Unix(),
//	})
//
//	jwtToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
//	return jwtToken
//}

// Logout Выход из аккаунта
//func Logout(c *gin.Context) {
//
//	token := models.Token{}
//
//	refreshToken, err := c.Cookie("refresh_token")
//
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
//		return
//	}
//
//	if err := models.DB.Where("refresh=?", refreshToken).First(&token).Error; err == nil {
//		models.DB.Delete(&token)
//	}
//
//	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
//	c.JSON(http.StatusOK, gin.H{"msg": "Вы успешно вышли из системы"})
//}

// CheckToken Проверка JWT
func CheckToken(token string) (uint, bool) {
	type MyCustomClaims struct {
		ID    uint   `json:"userid"`
		Email string `json:"email"`
		jwt.StandardClaims
	}

	tokenParse, _ := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if _, ok := tokenParse.Claims.(*MyCustomClaims); ok && tokenParse.Valid {
		if err := models.DB.Where("id=?", tokenParse.Claims.(*MyCustomClaims).ID).First(&models.User{}).Error; err == nil {
			return tokenParse.Claims.(*MyCustomClaims).ID, true
		} else {
			return 0, false
		}
	}

	return 0, false
}

// Refresh Создание JWT с помощью Refresh токена
//func Refresh(c *gin.Context) {
//	tokenRefresh, err := c.Cookie("refresh_token")
//	fmt.Println(tokenRefresh)
//	if err != nil {
//		c.JSON(http.StatusConflict, gin.H{"msg": "Рефреш токен не найден"})
//		return
//	}
//
//	token := models.Token{}
//
//	if err := models.DB.Where("refresh=?", tokenRefresh).First(&token).Error; err == nil && CheckToken(tokenRefresh) {
//		user := models.User{}
//
//		models.DB.Where("id=?", token.ID).First(&user)
//
//		newToken := CreateToken(user)
//		//хардкод - наше всё!
//		newTokenStruct := models.Token{Access: newToken}
//
//		models.DB.Model(&token).Updates(newTokenStruct)
//		c.JSON(http.StatusOK, gin.H{"access": newToken})
//		return
//	} else if !CheckToken(tokenRefresh) {
//		models.DB.Delete(&token)
//	}
//
//	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
//	c.JSON(http.StatusConflict, gin.H{"msg": "Рефреш токен не валидный"})
//}

func Activate(c *gin.Context) {
	emailCheck := models.EmailCheck{}
	if err := models.DB.Where("uuid=?", c.Param("uuid")).First(&emailCheck).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Подтверждение не существует или недействительно!"})
		return
	}

	user := models.User{}
	userUpdate := models.User{EmailCheck: true}

	if err := models.DB.Where("id=?", emailCheck.ID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Аккаунт не найден!"})
		models.DB.Delete(&emailCheck)
		return
	}

	models.DB.Model(&user).Updates(userUpdate)
	models.DB.Delete(&emailCheck)

	c.JSON(http.StatusOK, gin.H{"msg": "Аккаунт успешно активирован!"})
}

func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

func GetRole(c *gin.Context) {
	//token := c.Request.Header[""]
	c.Request.Header.Add("Authorization", "aboba")
	c.JSON(http.StatusOK, gin.H{"msg": "Проверь хэдер"})
}
