package services

import (
	"github.com/gin-gonic/gin"
	"go-server/models"
	"net/http"
)

func GetProfile(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id=?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Аккаунт не найден"})
		return
	}

	data := models.GetProfile{
		Email:           user.Email,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Role:            user.Role,
		Faculty:         user.Faculty,
		Group:           user.Group,
		FormOfEducation: user.FormOfEducation,
		Level:           user.Level,
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
