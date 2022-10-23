package services

import (
	"github.com/gin-gonic/gin"
	"go-server/models"
	"net/http"
)

func GetUsers(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}

	if _, role, ok := CheckToken(token, c); !ok || role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var users []models.User

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func DeleteUser(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}

	if _, role, ok := CheckToken(token, c); !ok || role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	if err := models.DB.Where("id=?", input.ID).First(&input).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Пользователь не найден"})
		return
	}

	models.DB.Delete(input)

	c.JSON(http.StatusOK, gin.H{"msg": "Успешно удалено"})
}

func ChangeRole(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}

	if _, role, ok := CheckToken(token, c); !ok || role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	models.DB.Model(&models.User{}).Where("id=?", input.ID).Update("role", input.Role)

	c.JSON(http.StatusOK, gin.H{"msg": "Успешно изменено на " + input.Role})
}

func ChangeData(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}

	if _, role, ok := CheckToken(token, c); !ok || role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var input, user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	if err := models.DB.Where("id=?", input.ID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Пользователь не найден"})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"msg": "Успешно изменено"})
}
