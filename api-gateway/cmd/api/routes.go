package main

import (
	"bet-manager/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authentication)
	server.POST("/login", handlers.Login)
	server.POST("/register", handlers.Register)
	authenticated.DELETE("/remove-user", handlers.RemoveUser)
	authenticated.GET("/get-all-users", handlers.GetAllUsers)
}
