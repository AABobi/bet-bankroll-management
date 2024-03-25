package main

import (
	"bet-manager/cmd/api/controllers"
	"bet-manager/repository"
	"bet-manager/server"
	"bet-manager/services"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB
var userRepository repository.IUserRepository

var handlers controllers.ApiHandlers

func main() {
	app := server.NewApp()

	repo, err := repository.NewPostgresUserRepository(app.Db)

	if err != nil {
		log.Fatalln("Problem with repository")
	}

	userService := services.NewUserService(repo)
	handlers = controllers.NewApiHandlersRepository(userService)
	RegisterRoutes(app.HttpServer)
	app.HttpServer.Run(":8080") //localhost:8080
}
