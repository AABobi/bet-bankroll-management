package services

import (
	"bet-manager/models"
	"bet-manager/repository"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserService interface {
	Login(context *gin.Context) (uint, []byte)
	Register(context *gin.Context) (uint, []byte)
	RemoveUser(context *gin.Context) (uint, []byte)
	GetAllUsers(context *gin.Context) (uint, []byte)
}

type userService struct {
	Repository repository.IUserRepository
}

type serviceMessage struct {
	message string
}

func NewUserService(repo repository.IUserRepository) UserService {
	return &userService{repo}
}

func (api *userService) Login(context *gin.Context) (uint, []byte) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		messageString := `{"message": "Could not parse request data"}`
		return marshalAndStatusHandler(messageString, http.StatusBadRequest)
	}

	dbUser, err := api.Repository.GetUser(user.Email)

	if err != nil {
		messageString := `{"message": "Could not get user"}`
		return marshalAndStatusHandler(messageString, http.StatusBadRequest)
	}

	dbUserByte, err := json.Marshal(&dbUser)

	return marshalAndStatusHandler(dbUserByte, http.StatusBadRequest)

	/*	token, err := auth.GenerateToken(user.Email, dbUser.UserId)

		if err != nil {
			messageString := `{"Problem with generate token"}`
			a, err := json.Marshal(messageString)

			if err != nil {
				log.Panicln("Cannot marshal message")
			}
			return 401, a
		}

		return 200,*/
	//context.Request.Header.Set("token", token)
	//context.Header("Authorization", token)
	//context.JSON(http.StatusOK, dbUser)
}

func (api *userService) Register(context *gin.Context) (uint, []byte) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {

		//context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		message := `{"message": "Could not parse data"}`
		return marshalAndStatusHandler(message, http.StatusBadRequest)
	}
	existUser, err := api.Repository.GetUser(user.Email)

	if existUser.Email == user.Email && err == nil {
		//context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot create user because user with that email exists"})
		message := `{"message": "Cannot create user because user with that email exists"}`
		return marshalAndStatusHandler(message, http.StatusBadRequest)
	} else {
		err = api.Repository.CreateUser(user)
		if err != nil {
			//context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot create user"})
			message := `{"message": "Cannot create user"}`
			return marshalAndStatusHandler(message, http.StatusBadRequest)
		}

		//context.JSON(http.StatusCreated, gin.H{"message": "User created"})
		message := `{"message": "User created"}`
		return marshalAndStatusHandler(message, http.StatusCreated)
	}
}

func (api *userService) RemoveUser(context *gin.Context) (uint, []byte) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		//context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Could parse data"})
		message := `{"message": "Could parse data"}`
		return marshalAndStatusHandler(message, http.StatusBadRequest)
	}
	existUser, err := api.Repository.GetUser(user.Email)
	if err != nil {
		//context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Cannot remove user because it doesn't exist"})
		message := `{"message": "Cannot remove user because it doesn't exist"}`
		return marshalAndStatusHandler(message, http.StatusBadRequest)
	}
	if existUser.Email == user.Email {
		err = api.Repository.RemoveUser(user)
		if err != nil {
			//context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("%v", err)})
			message := fmt.Sprintf(`{"message": "%v"}`, err)
			return marshalAndStatusHandler(message, http.StatusBadRequest)
		}
		message := `{"message": "User removed"}`
		return marshalAndStatusHandler(message, http.StatusOK)
		//context.JSON(http.StatusOK, gin.H{"message": "User removed"})
	}
	message := `{"message": "Could parse data"}`
	return marshalAndStatusHandler(message, http.StatusBadRequest)
}

func (api *userService) GetAllUsers(context *gin.Context) (uint, []byte) {
	var users []models.User
	users, err := api.Repository.GetAll()
	if err != nil {
		//context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Cannot get users"})
		message := `{"message": "Cannot get users"}`
		return marshalAndStatusHandler(message, http.StatusBadRequest)
	}
	//message := `{"message": "Cannot get users"}`
	return marshalAndStatusHandler(users, http.StatusOK)
	//context.JSON(http.StatusOK, users)
}

func marshalAndStatusHandler(message any, status uint) (uint, []byte) {
	messageByte, err := json.Marshal(message)
	if err != nil {
		errorMessage := `{"message": "Could not marshal message"}`
		log.Println("Could not marshal message")
		errorMessageByte, err := json.Marshal(errorMessage)

		if err != nil {
			log.Panicln("Could not marshal message")
		}

		return 500, errorMessageByte
	}
	return status, messageByte
}
