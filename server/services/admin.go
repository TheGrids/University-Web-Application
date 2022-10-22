package services

import (
	"github.com/gin-gonic/gin"
	"go-server/models"
	"net/http"
)

func GetUsers(c *gin.Context) {
	if _, role, ok := CheckToken(c.Request.Header.Get("Authorization"), c); !ok || role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Недостаточно прав."})
		return
	}

	var users []models.User

	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
