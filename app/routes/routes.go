package routes

import (
	"ap/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//server.GET("/getUsers", getUsers)
	server.POST("/signup", singup)
	server.POST("login", login)
	server.GET("/dockerTest", dockerTest)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/getUsers", getUsers)
	authenticated.GET("/getBets", getBets)
	authenticated.POST("/addBet", addBet)
}
