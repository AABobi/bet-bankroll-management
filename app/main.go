package main

import (
	"app/db"
	"app/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	fmt.Println("test")
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080
}
