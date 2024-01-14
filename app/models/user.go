package models

import (
	"app/db"
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
	/*query := "INSERT INTO users(login, email, password) VALUES (? , ? , ?)"

	stmt, err := db.DB.Prepare(query)
	fmt.Println("SaveUser1", u.Login, u.Email, u.Password)
	if err != nil {
		return err
	}

	defer stmt.Close()
	fmt.Println("SaveUser2", u.Email)
	_, err = stmt.Exec(u.Login, u.Email, u.Password)

	if err != nil {
		return err
	}

	return nil*/
	query := "INSERT INTO users(login, email, password) VALUES (?, ?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("save1")
		return err
	}

	defer stmt.Close()

	//Hash password
	fmt.Println(u.Password)
	//hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		fmt.Println("save2")
		return err
	}

	result, err := stmt.Exec(u.Login, u.Email, u.Password)

	if err != nil {
		fmt.Println("save3")
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	fmt.Println("save4")
	return err
}
