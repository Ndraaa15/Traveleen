package middleware

import (
	"gin/sdk/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsUserLoggedIn(ctx *gin.Context) {
	header := ctx.Request.Header.Get("Authorization")

	if header == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You must be logged in first."})
		return
	}

	tokenParts := strings.SplitN(header, " ", 2)
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	tokenString := tokenParts[1]

	tokenClaims, err := jwt.DecodeJWTToken(tokenString)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid JWT token"})
	}

	userID := uint(tokenClaims["id"].(float64))

	ctx.Set("user", userID)
	ctx.Next()
}

func CORS(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Length, Content-Type, Authorization")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}
