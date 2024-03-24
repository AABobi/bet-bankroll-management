package middlewares

import (
	"bet-manager/auth"
	"context"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authentication(ginCtx *gin.Context) {
	token := ginCtx.Request.Header.Get("Authorization")

	parts := strings.Split(token, " ")

	if len(parts) >= 2 && parts[0] == "Bearer" {
		// Extract the token
		token = parts[1]
	} else {
		ginCtx.AbortWithStatusJSON(401, gin.H{"message": "No access token"})
		return
	}

	userId, err := auth.VerifyToken(token)
	if err != nil {
		ginCtx.AbortWithStatusJSON(401, gin.H{"message": "Incorrect access token"})
		return
	}

	//ginCtx.Set("userId", userId)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", userId)

	ginCtx.Next()
}
