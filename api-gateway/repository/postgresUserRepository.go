package repository

import (
	"bet-manager/models"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) (*PostgresUserRepository, error) {
	if db == nil {
		return nil, errors.New("database cannot be nil")
	}

	return &PostgresUserRepository{db: db}, nil
}

func (repo *PostgresUserRepository) GetUser(email string) (models.User, error) {
	fmt.Println(email)
	query := `SELECT * FROM users WHERE email = $1`

	row := repo.db.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.UserId, &user.Email, &user.Password, &user.Role)

	if err != nil {
		return models.User{}, errors.New("Cannot find user")
	}

	return user, nil
}

func (repo *PostgresUserRepository) CreateUser(user models.User) error {
	query := `INSERT INTO users(email, password, role) 
    values($1,$2,$3)
`

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Email, user.Password, user.Role)
	if err != nil {
		return errors.New("Cannot create new user")
	}

	defer stmt.Close()
	return nil
}

func (repo *PostgresUserRepository) RemoveUser(user models.User) error {
	query := `DELETE FROM users WHERE email = $1`

	stmt, err := repo.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Email)

	if err != nil {
		return errors.New("Cannot remove user")
	}
	defer stmt.Close()
	return nil
}

func (repo *PostgresUserRepository) GetAll() ([]models.User, error) {
	var users = []models.User{}
	query := `SELECT * FROM users`

	rows, err := repo.db.Queryx(query)

	if err != nil {
		return []models.User{}, errors.New("Cannot proceed get all users")
	}
	user := models.User{}
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}
	defer rows.Close()
	return users, nil
}
