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
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if err := models.DB.Where("id=?", input.ID).Delete; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Успешно удалено"})
}
