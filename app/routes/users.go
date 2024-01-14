package routes

//"fmt"
//	"net/http"
//	"app/db"
import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Login    string
	Password string `binding:"required"`
}

func getUsers(context *gin.Context) {
	resposne, err := models.GetAllUser()

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Cannot retrieve users"})
	}

	context.JSON(http.StatusOK, resposne)
}

func singup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	fmt.Print("test11")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	}

	err = user.SaveUser()
	fmt.Print("test22")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot signup user"})
		return
	}
	fmt.Print("test33")
	context.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}
