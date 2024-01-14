package routes

import (
	"app/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//server.GET("/getUsers", getUsers)
	server.POST("/signup", singup)
	server.POST("login", login)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/getUsers", getUsers)
}
