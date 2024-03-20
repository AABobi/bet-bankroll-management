package repository

import "bet-manager/models"

type IUserRepository interface {
	GetAll() ([]models.User, error)
	GetUser(email string) (models.User, error)
	CreateUser(user models.User) error
	RemoveUser(user models.User) error
}
