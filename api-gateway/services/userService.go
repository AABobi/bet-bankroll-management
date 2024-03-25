package services

import (
	"bet-manager/auth"
	"bet-manager/models"
	"bet-manager/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserService interface {
	Login(context *gin.Context) (int, *models.User, *gin.H, *string)
	Register(context *gin.Context) (int, gin.H)
	RemoveUser(context *gin.Context) (int, gin.H)
	GetAllUsers(context *gin.Context) (int, *[]models.User, *gin.H)
}

type userService struct {
	Repository repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) UserService {
	return &userService{repo}
}

func (api *userService) Login(ctx *gin.Context) (int, *models.User, *gin.H, *string) {
	var user models.User
	fmt.Println("1")
	err := ctx.ShouldBindJSON(&user)
	fmt.Println("11")
	if err != nil {
		fmt.Println("2")
		ginMessage := gin.H{"message": "Could not parse request data"}
		return http.StatusBadRequest, nil, &ginMessage, nil
	}

	fmt.Println("22")
	dbUser, err := api.Repository.GetUser(user.Email)
	fmt.Println("222")
	if err != nil {
		fmt.Println("3")
		ginMessage := gin.H{"message": "Could not find user"}
		return http.StatusBadRequest, nil, &ginMessage, nil
	}

	fmt.Println("4")
	token, _ := auth.GenerateToken(dbUser.Email, dbUser.UserId)
	return http.StatusOK, &dbUser, nil, &token
}

func (api *userService) Register(context *gin.Context) (int, gin.H) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	fmt.Println("user")
	fmt.Println(user)
	fmt.Println("0")
	if err != nil {
		fmt.Println("1")
		ginMessage := gin.H{"message": "Could not parse data"}
		return http.StatusBadRequest, ginMessage
	}
	fmt.Println("11")
	existUser, err := api.Repository.GetUser(user.Email)
	fmt.Println("111")
	if existUser.Email == user.Email && err == nil {
		fmt.Println("2")
		ginMessage := gin.H{"message": "Cannot create user because user with that email exists"}
		return http.StatusBadRequest, ginMessage
	} else {
		fmt.Println("3")
		err = api.Repository.CreateUser(user)
		if err != nil {
			fmt.Println("4")
			ginMessage := gin.H{"message": "Cannot create user"}
			return http.StatusBadRequest, ginMessage
		}
		fmt.Println("5")
		ginMessage := gin.H{"message": "User created"}
		return http.StatusCreated, ginMessage
	}
}

func (api *userService) RemoveUser(context *gin.Context) (int, gin.H) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		ginMessage := gin.H{"message": "Could parse data"}
		return http.StatusBadRequest, ginMessage
	}
	existUser, err := api.Repository.GetUser(user.Email)
	if err != nil {
		ginMessage := gin.H{"message": "Cannot remove user because it doesn't exist"}
		return http.StatusBadRequest, ginMessage
	}
	if existUser.Email == user.Email {
		err = api.Repository.RemoveUser(user)
		if err != nil {
			ginMessage := gin.H{"message": fmt.Sprintf("%v", err)}
			return http.StatusBadRequest, ginMessage
		}
		ginMessage := gin.H{"message": "User removed"}
		return http.StatusOK, ginMessage
	}
	ginMessage := gin.H{"message": "Could parse data"}
	return http.StatusBadRequest, ginMessage
}

func (api *userService) GetAllUsers(context *gin.Context) (int, *[]models.User, *gin.H) {
	var users []models.User
	users, err := api.Repository.GetAll()
	fmt.Println("gellALl")
	fmt.Println(users)
	if err != nil {
		ginMessage := gin.H{"message": "Cannot get users"}
		return http.StatusBadRequest, nil, &ginMessage
	}
	fmt.Println("ALE GIT?")
	return http.StatusOK, &users, nil
}
