package controllers

import (
	"bet-manager/services"
	"github.com/gin-gonic/gin"
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

func (api *apiHandlers) Login(ctx *gin.Context) {
	status, user, data, token := api.Service.Login(ctx)
	if status == 200 {
		ctx.Header("Authorization", *token)
		ctx.JSON(status, user)
	} else {
		ctx.AbortWithStatusJSON(status, data)
	}
}

func (api *apiHandlers) Register(ctx *gin.Context) {
	status, data := api.Service.Register(ctx)
	if status == 201 {
		ctx.JSON(status, data)
	} else {
		ctx.AbortWithStatusJSON(status, data)
	}
}

func (api *apiHandlers) RemoveUser(ctx *gin.Context) {
	status, data := api.Service.RemoveUser(ctx)
	if status == 200 {
		ctx.JSON(status, data)
	} else {
		ctx.AbortWithStatusJSON(status, data)
	}
}

func (api *apiHandlers) GetAllUsers(ctx *gin.Context) {
	status, users, data := api.Service.GetAllUsers(ctx)
	if status == 200 {
		ctx.JSON(status, users)
	} else {
		ctx.AbortWithStatusJSON(status, data)
	}
}
