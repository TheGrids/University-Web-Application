package services

import (
	"github.com/gin-gonic/gin"
	"go-server/models"
	"net/http"
	"time"
)

func AddMessage(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}

	id, _, ok := CheckToken(token, c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var input models.AddMessage

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	message := models.Message{AuthorId: id, Title: input.Title, Body: input.Body, Time: time.Now().Unix()}

	models.DB.Create(&message)

	c.JSON(http.StatusOK, gin.H{"msg": "Успешно создано"})
}

func GetMessages(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}

	if _, role, ok := CheckToken(token, c); !ok || (role != "admin") {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var messages []models.Message

	models.DB.Order("time DESC, id DESC").Find(&messages)

	c.JSON(http.StatusOK, gin.H{"data": messages})
}

func DeleteMessage(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}

	if _, role, ok := CheckToken(token, c); !ok || role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var input models.Message

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	if err := models.DB.Where("id=?", input.ID).First(&input).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Сообщение не найдено"})
		return
	}

	models.DB.Delete(input)

	c.JSON(http.StatusOK, gin.H{"msg": "Успешно удалено"})
}
