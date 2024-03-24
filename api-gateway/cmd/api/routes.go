package main

import (
	"bet-manager/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authentication)
	server.POST("/login", api.Login)
	server.POST("/register", api.Register)
	authenticated.DELETE("/remove-user", api.RemoveUser)
	authenticated.GET("/get-all-users", api.GetAllUsers)
}
