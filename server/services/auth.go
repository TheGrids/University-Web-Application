package services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-server/models"
	"net/http"
)

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

	user := models.User{Email: input.Email, Password: input.Password, FirstName: input.FirstName, LastName: input.LastName, EmailCheck: false}
	models.DB.Create(&user)

	emailCheck := models.EmailCheck{ID: user.ID, UUID: uuid.New().String()}
	models.DB.Create(&emailCheck)

	SendEmailUUID(user.Email, emailCheck.UUID)

	c.JSON(http.StatusOK, gin.H{"msg": "Мы отправили письмо с подтверждением на вашу электронную почту. Вы сможете зайти на аккант только после подтверждения."})
}
