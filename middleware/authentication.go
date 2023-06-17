package middleware

import (
	"finalproject2/helpers"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)
		_ = verifyToken

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid email or password",
			})
			log.Println(err.Error())
			return
		}

		ctx.Set("userID", int(verifyToken.(jwt.MapClaims)["id"].(float64)))
		ctx.Set("email", verifyToken.(jwt.MapClaims)["email"].(string))
		// ctx.Set("userData", verifyToken)
		ctx.Next()

	}
}
