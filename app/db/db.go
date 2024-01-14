package db

//_ oznacz, zeby autosave nie usuną
import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	DB = db

	createTables()
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	fmt.Println("Tables created successfully!")
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		login TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table")
	}

	createBetsTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userID INTEGER NOT NULL,
		betValue INTEGER,
        data DATETIME,
		tax INTEGER,
		win BOOLEAN , 
		FOREIGN KEY(userID) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createBetsTable)

	if err != nil {
		panic("Could not create users table")
	}
}
