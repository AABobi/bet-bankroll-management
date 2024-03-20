package main

import (
	"bet-manager/repository"
	"bet-manager/server"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB
var userRepository repository.IUserRepository

func main() {
	app := server.NewApp()

	repo, err := repository.NewPostgresUserRepository(app.Db)

	if err != nil {
		log.Fatalln("Problem with repository")
	}

	userRepository = repo

	RegisterRoutes(app.HttpServer)
	app.HttpServer.Run(":8080") //localhost:8080
}
