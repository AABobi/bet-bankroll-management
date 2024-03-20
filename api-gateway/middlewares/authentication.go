package middlewares

import (
	"bet-manager/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	// Split the string by space
	parts := strings.Split(token, " ")

	// Check if there are at least two parts and the first part is "Bearer"
	if len(parts) >= 2 && parts[0] == "Bearer" {
		// Extract the token
		token = parts[1]
	} else {
		context.AbortWithStatusJSON(401, gin.H{"message": "No access token"})
		return
	}

	userId, err := auth.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{"message": "Incorrect access token"})
		return
	}

	context.Set("userId", userId)
	fmt.Println(userId)
	context.Next()
}
