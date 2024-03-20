package server

import (
	"bet-manager/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type App struct {
	HttpServer *gin.Engine
	Db         *sqlx.DB
}

func NewApp() *App {
	db := initDB()
	server := gin.Default()

	return &App{
		HttpServer: server,
		Db:         db,
	}
}
func initDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable password=mysecretpassword host=localhost")

	if err != nil {
		fmt.Println("Test")
		log.Fatalln(err)
		fmt.Println("Test")
	}
	//defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfuly connected")
	}

	_, err = db.Exec(`
	           CREATE TABLE IF NOT EXISTS users (
	               UserId SERIAL PRIMARY KEY ,
	               Email VARCHAR(100),
	               Password VARCHAR(100),
	               Role INT
	           )
	       `)
	if err != nil {
		panic(err)
	}
	fmt.Println("Users table created successfully")

	// Iterate over the rows and scan into user structs
	/*	users1 := []models.User{
			{Email: "test", Password: "test", Role: 1},
		}
		query1 := `INSERT INTO users(email, password, role) VALUES ($1, $2, $3)`
		stmt, err := db.Prepare(query1)

		if err != nil {
			fmt.Println("save1")
		}
		defer stmt.Close()

		result, err := stmt.Exec(users1[0].Email, users1[0].Password, 1)
		userId, err := result.LastInsertId()

		fmt.Println(userId)
		fmt.Println("save4")*/

	query := `SELECT * FROM users`
	// Execute the query
	rows, err := db.Queryx(query)

	defer rows.Close()
	var users []models.User
	for rows.Next() {
		fmt.Println("ROWS NEXT")
		var user models.User
		if err := rows.StructScan(&user); err != nil {
			fmt.Println("EROR")
			fmt.Println(err)
		}
		users = append(users, user)
	}
	fmt.Println(len(users))
	fmt.Println(users)
	fmt.Println(users[0])
	/*
	 */
	return db
}

/*package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

var DB *sqlx.DB
var usersTest = []User{}

func dbInit() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable password=mysecretpassword host=postgres")

	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfuly connected")
	}

	_, err = db.Exec(`
	           CREATE TABLE IF NOT EXISTS users (
	               username VARCHAR(100),
	               email VARCHAR(100)
	           )
	       `)
	if err != nil {
		panic(err)
	}
	fmt.Println("Users table created successfully")

	// Insert 5 users into the table
	users := []User{
		{"user1", "user1@example.com"},
		{"user2", "user2@example.com"},
		{"user3", "user3@example.com"},
		{"user4", "user4@example.com"},
		{"user5", "user5@example.com"},
	}

	for _, user := range users {
		_, err := db.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", user.Name, user.Email)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Inserted user: %s\n", user.Name)
	}

	place := User{}
	rows, _ := db.Queryx("SELECT username, email FROM users")
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%#v\n", place)
		usersTest = append(usersTest, place)
	}

	DB = db
}

func test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, usersTest)
}
*/
