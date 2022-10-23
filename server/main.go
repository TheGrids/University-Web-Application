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
		AllowOrigins:     []string{"http://localhost:3000", "https://universityweb.site", "http://127.0.0.1:3000"},
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
		api.POST("/addnews", services.AddNews)
		api.GET("/news", services.GetNews)
		api.GET("/news/:id", services.GetNew)
		api.DELETE("/deletenews", services.DeleteNews)
		api.POST("/newssort", services.NewsSort)
		api.POST("/addmessage", services.AddMessage)
		api.GET("/messages", services.GetMessages)
		api.DELETE("/deletemessage", services.DeleteMessage)

		admin := api.Group("/admin")
		{
			admin.GET("/profiles", services.GetUsers)
			admin.PUT("/changerole", services.ChangeRole)
			admin.PUT("/changedata", services.ChangeData)
			admin.DELETE("/person/delete", services.DeleteUser)
		}
	}

	r.Run()
}
