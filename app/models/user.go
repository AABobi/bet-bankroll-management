package models

import (
	"ap/db"
	"ap/utils"

	"errors"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Login    string
	Password string `binding:"required"`
}

func GetAllUser() ([]User, error) {
	query := "SELECT * from users"

	correctRows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer correctRows.Close()

	var users = []User{}

	for correctRows.Next() {
		var user User

		err = correctRows.Scan(&user.ID, &user.Email, &user.Login, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	fmt.Println("testGetALlUsers")
	return users, nil
}

func (u User) SaveUser() error {
	query := "INSERT INTO users(login, email, password) VALUES (? , ? , ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	_, err = stmt.Exec(u.Login, u.Email, hashedPassword)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) CheckCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	singleRow := db.DB.QueryRow(query, u.Email)

	var password string

	err := singleRow.Scan(&u.ID, &password)

	if err != nil {
		fmt.Println("test")
		return err
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, password)

	if !isPasswordValid {
		fmt.Println("test1")
		return errors.New("Invalid credentials")
	}

	return nil
}
