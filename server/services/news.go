package services

import (
	"github.com/gin-gonic/gin"
	"go-server/models"
	"net/http"
	"time"
)

func AddNews(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}

	id, role, ok := CheckToken(token, c)
	if !ok || (role != "admin" && role != "teacher") {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var input models.AddNews

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	news := models.News{AuthorId: id, Title: input.Title, Body: input.Body, Tag: input.Tag, Time: time.Now().Unix()}

	if err := models.DB.Where("Title=?", input.Title).First(&models.News{}).Error; err == nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Статья с таким заголовком уже есть."})
		return
	}

	models.DB.Create(&news)

	c.JSON(http.StatusOK, gin.H{"msg": "Успешно создано"})
}

func GetNews(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}

	if _, role, ok := CheckToken(token, c); !ok || (role != "admin" && role != "teacher") {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var news []models.News

	models.DB.Order("time DESC, id DESC").Find(&news)

	c.JSON(http.StatusOK, gin.H{"data": news})
}

func GetNew(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}

	if _, role, ok := CheckToken(token, c); !ok || (role != "admin" && role != "teacher") {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var news models.News
	if err := models.DB.Where("id=?", c.Param("id")).First(&news).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Новость не найдена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": news})
}

func DeleteNews(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}

	if _, role, ok := CheckToken(token, c); !ok || role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var input models.News

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	if err := models.DB.Where("id=?", input.ID).First(&input).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Новость не найдена"})
		return
	}

	models.DB.Delete(input)

	c.JSON(http.StatusOK, gin.H{"msg": "Успешно удалено"})
}

func NewsSort(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не найден токен."})
		return
	}

	if _, role, ok := CheckToken(token, c); !ok || role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var input models.News

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	var news []models.News

	if input.Tag == "Учебные новости" || input.Tag == "Социальная жизнь" || input.Tag == "Жизнь ВУЗа" {
		models.DB.Where("tag", input.Tag).Order("time DESC, id DESC").Find(&news)
		c.JSON(http.StatusOK, gin.H{"data": news})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"msg": "Отсутствуют новости с такой меткой"})
}
