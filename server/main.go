package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-server/models"
	"go-server/services"
	"time"
)

func main() {
	r := gin.Default()

	models.ConnectionDataBase()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://universityweb.site"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           1 * time.Minute,
	}))

	api := r.Group("/api")
	{
		api.POST("/register", services.RegisterUser)
		api.GET("/activate/:uuid", services.Activate)
		api.POST("/logout", services.Logout)
		api.POST("/login", services.LoginUser)
		api.POST("/register", services.RegisterUser)
		api.POST("/refresh", services.Refresh)
	}

	r.Run()
}
