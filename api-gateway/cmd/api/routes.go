package main

import (
	"bet-manager/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authentication)
	server.POST("/login", login)
	server.POST("/register", register)
	authenticated.DELETE("/remove-user", removeUser)
	authenticated.GET("/get-all-users", getAllUsers)
}
