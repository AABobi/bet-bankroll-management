package mocks

import (
	"bet-manager/models"
	"fmt"
	"github.com/stretchr/testify/mock"
)

type IUserRepository struct {
	mock.Mock
}

func (_m *IUserRepository) GetAll() ([]models.User, error) {
	ret := _m.Called()
	fmt.Println("1")
	var r0 []models.User
	fmt.Println("2")
	if rf, ok := ret.Get(0).(func() []models.User); ok {
		fmt.Println("3")
		r0 = rf()
	} else {
		fmt.Println("4")
		if ret.Get(0) != nil {
			fmt.Println("5")
			r0 = ret.Get(0).([]models.User)
		}
	}
	fmt.Println("6")
	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}
	fmt.Println("1")
	return r0, r1
}

func (_m *IUserRepository) GetUser(email string) (models.User, error) {
	ret := _m.Called(email)

	var r0 models.User
	if rf, ok := ret.Get(0).(func() models.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mock CreateUser method
func (_m *IUserRepository) CreateUser(user models.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Mock RemoveUser method
func (_m *IUserRepository) RemoveUser(user models.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
