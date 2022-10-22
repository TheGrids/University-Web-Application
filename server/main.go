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
		AllowOrigins:     []string{"*localhost:3000", "https://universityweb.site"},
		AllowHeaders:     []string{"Role", "Origin", "Content-Type", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
		MaxAge:           1 * time.Minute,
	}))

	api := r.Group("/api")
	{
		api.POST("/register", services.RegisterUser)
		api.GET("/activate/:uuid", services.Activate)
		//api.POST("/logout", services.Logout)
		api.POST("/login", services.LoginUser)
		//api.POST("/refresh", services.Refresh)
		api.GET("/profile/:id", services.GetProfile)
		api.GET("/verification", services.Verification)
	}

	r.Run()
}
