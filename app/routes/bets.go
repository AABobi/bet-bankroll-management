package routes

import (
	"ap/models"
	//	"app/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func dockerTest(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "DockerTest"})
}

func addBet(context *gin.Context) {
	var bet models.Bet

	err := context.ShouldBindJSON(&bet)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot add new bet"})
	}
	fmt.Println("bet")
	fmt.Println(bet)
	fmt.Println(&bet)
	err = bet.SaveBet()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot save new bet"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Bet added", "bet": bet})
}

func getBets(context *gin.Context) {
	bets, err := models.GetAllBets()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Cannot retrieve bets"})
	}

	context.JSON(http.StatusOK, bets)
}
