package controllers

import (
	"bet-manager/models"
	"bet-manager/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiHandlers interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
	RemoveUser(context *gin.Context)
	GetAllUsers(context *gin.Context)
}

type apiHandlers struct {
	Service services.UserService
}

func NewApiHandlersRepository(service services.UserService) ApiHandlers {
	return &apiHandlers{
		Service: service,
	}
}

func (api *apiHandlers) Login(context *gin.Context) {
	api.Re
	/*var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": "Could not parse request data"})
		return
	}

	dbUser, err := api.Repository.GetUser(user.Email)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Could not get user"})
		return
	}

	token, err := auth.GenerateToken(user.Email, dbUser.UserId)
	if err != nil {
		context.AbortWithStatusJSON(401, "Problem with generate token")
		return
	}
	//context.Request.Header.Set("token", token)
	context.Header("Authorization", token)
	context.JSON(http.StatusOK, dbUser)*/
}

func (api *apiHandlers) Register(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	existUser, err := api.Repository.GetUser(user.Email)

	if existUser.Email == user.Email && err == nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot create user because user with that email exists"})
		return
	} else {
		err = api.Repository.CreateUser(user)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot create user"})
			return
		}

		context.JSON(http.StatusCreated, gin.H{"message": "User created"})
	}
}

func (api *apiHandlers) RemoveUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could parse data"})
		return
	}
	existUser, err := api.Repository.GetUser(user.Email)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Cannot remove user because it doesn't exist"})
		return
	}
	if existUser.Email == user.Email {
		err = api.Repository.RemoveUser(user)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "User removed"})
	}
}

func (api *apiHandlers) GetAllUsers(context *gin.Context) {
	var users []models.User
	users, err := api.Repository.GetAll()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot get users"})
		return
	}

	context.JSON(http.StatusOK, users)
}
