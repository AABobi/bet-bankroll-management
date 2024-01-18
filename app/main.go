package main

import (
	"ap/db"
	"ap/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	/*db.InitDB()
	fmt.Println("test")
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":9000") //localhost:8080*/
	server := gin.Default()

	// Apply CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://frontend-vue2-service:80"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}

	server.Use(cors.New(config))
	//gin.Default().SetTrustedProxies([]string{"localhost:3000", ":3000"})
	// Register your routes
	routes.RegisterRoutes(server)

	// Start the server
	server.Run()
}
